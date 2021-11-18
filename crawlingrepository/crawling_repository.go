package crawlingrepository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
	"upsider.crawling/config"
	"upsider.crawling/crawlingproto"
)

type CrawlingInterface interface {
	Crawling(pass string, input *crawlingproto.UserInput) ([]*User, []*Bank, []*Detail, error)
}

type crawlingRepository struct{}

func NewCrawling() CrawlingInterface {
	return &crawlingRepository{}
}

type User struct {
	officeName string
	UserId     string
	UpdatedAt  time.Time
}

type Bank struct {
	Id              string
	UserId          string
	BankId          string
	OfficeName      string
	BankName        string
	Amount          int64
	LastCommit      string
	ResitrationDate string
	BankStatus      string
	Kind            string
}

type Detail struct {
	Id                string
	UserId            string
	BankId            string
	TradingDate       string
	TradingContent    string
	Payment           int64
	Amount            int64
	BankName          string
	Status            string
	TransactionNumber string
	edit              string
	Crawling          time.Time
}

type DetailUrl struct {
	DetailUrl string
}

func (c *crawlingRepository) Crawling(pass string, input *crawlingproto.UserInput) ([]*User, []*Bank, []*Detail, error) {
	var user []*User
	var banks []*Bank
	var details []*Detail
	var detailUrls []*DetailUrl

	ctx, cancel := context.WithTimeout(config.NewChromedpContext(), 3*time.Minute)
	defer cancel()

	loginURL := "https://erp.moneyforward.com/session/new"
	topURL := "https://erp.moneyforward.com/home"
	illegalCheck := ""

	loginIdSel := `/html/body/main/div/div/div/div/div[2]/div/div[2]/div[1]/section/form/div[2]/div/input`
	loginPassSel := `/html/body/main/div/div/div/div/div[2]/div/div[2]/div[1]/section/form/div[2]/div/input[2]`
	nextLoginButtonSel := `/html/body/main/div/div/div/div/div[2]/div/div[2]/div[1]/section/form/div[2]/div/div[3]/input`

	loginActionFunc := chromedp.ActionFunc(func(ctx context.Context) error {
		chromedp.Navigate(loginURL).Do(ctx)
		chromedp.Location(&illegalCheck).Do(ctx)
		if illegalCheck == "chrome-error://chromewebdata/" {
			return fmt.Errorf("URLの遷移に失敗しました: %s", illegalCheck)
		}
		chromedp.SendKeys(loginIdSel, input.UserId, chromedp.NodeVisible).Do(ctx)
		chromedp.WaitVisible(`._17NyHEtk.submitBtn._2JooeuSw.bizDomain`, chromedp.NodeVisible).Do(ctx)
		chromedp.Click(`._17NyHEtk.submitBtn._2JooeuSw.bizDomain`, chromedp.BySearch).Do(ctx)

		chromedp.WaitVisible(`._1SLW-zrQ.inputItem`, chromedp.ByQuery).Do(ctx)
		chromedp.SendKeys(loginPassSel, pass, chromedp.BySearch).Do(ctx)
		chromedp.Click(nextLoginButtonSel).Do(ctx)
		chromedp.WaitVisible(`body`, chromedp.NodeVisible).Do(ctx)

		chromedp.Location(&illegalCheck).Do(ctx)
		if illegalCheck != topURL {
			return fmt.Errorf("ログインに失敗しました: %s", illegalCheck)
		}

		return nil
	})

	crawlingActionFunc := chromedp.ActionFunc(func(ctx context.Context) error {

		chromedp.Click(`#accounting`, chromedp.NodeVisible).Do(ctx)
		chromedp.WaitVisible(`body`, chromedp.ByQuery).Do(ctx)
		chromedp.Location(&illegalCheck).Do(ctx)
		if !strings.Contains(illegalCheck, "https://accounting.moneyforward.com/home") {
			return fmt.Errorf("確定申告のトップページに遷移できませんでした: %s", illegalCheck)
		}
		chromedp.WaitVisible(`body`, chromedp.ByQuery).Do(ctx)
		officeNodes := []*cdp.Node{}

		chromedp.Nodes(`#dropdown-office`, &officeNodes, chromedp.ByQuery).Do(ctx)
		if len(officeNodes) == 0 {
			return fmt.Errorf("会社名のNodeが取得できませんでした。")
		}
		res, err := dom.GetOuterHTML().WithNodeID(officeNodes[0].NodeID).Do(ctx)
		if err != nil {
			fmt.Printf("会社名のdomが取得できませんでした %s", err)
		}

		user, err = scrapingOfOfficeName(res, user)
		if err != nil {
			fmt.Printf("会社名が取得できませんでした %s", err)
		}

		chromedp.Location(&illegalCheck).Do(ctx)
		account := illegalCheck[40:]
		registeredListUrl := "https://accounting.moneyforward.com/accounts" + account
		chromedp.Navigate(registeredListUrl).Do(ctx)
		chromedp.Location(&illegalCheck).Do(ctx)
		if !strings.Contains(illegalCheck, "https://accounting.moneyforward.com/accounts") {
			return fmt.Errorf("明細の登録済み一覧ページに遷移できませんでした: %s", illegalCheck)
		}
		chromedp.WaitVisible(`body`, chromedp.ByQuery).Do(ctx)

		fetchBanks := chromedp.ActionFunc(func(ctx context.Context) error {
			chromedp.WaitReady(`.text-center`, chromedp.ByQuery).Do(ctx)
			bankNodes := []*cdp.Node{}

			chromedp.Nodes(`.ca-table-horizontal.mf-mb10`, &bankNodes, chromedp.ByQueryAll).Do(ctx)
			if len(bankNodes) == 0 {
				return fmt.Errorf("銀行、並びにカード情報が取れませんでした。")
			}
			res, err := dom.GetOuterHTML().WithNodeID(bankNodes[0].NodeID).Do(ctx)
			if err != nil {
				fmt.Printf("error %s", err)
			}

			banks, detailUrls, err = scrapingOfBanks(res, banks, user, detailUrls)
			if err != nil {
				fmt.Printf("error %s", err)
			}

			for i, n := range detailUrls {
				kind, err := fetchBanksKind(ctx, n.DetailUrl, registeredListUrl)
				if err != nil {
					return fmt.Errorf("銀行、並びにカードの種別ができませんでした。")
				}
				banks[i].Kind = kind
			}
			return nil
		})

		fetchDetails := chromedp.ActionFunc(func(ctx context.Context) error {
			continuationNode := []*cdp.Node{}
			detailsNode := []*cdp.Node{}
			bankIdNode := []*cdp.Node{}
			for _, n := range detailUrls {
				chromedp.Navigate(n.DetailUrl).Do(ctx)
				chromedp.WaitVisible(`#js-acts-table-tbody`, chromedp.NodeVisible).Do(ctx)
				time.Sleep(3 * time.Second)
				chromedp.Location(&illegalCheck).Do(ctx)
				if !strings.Contains(illegalCheck, "https://accounting.moneyforward.com/accounts/trans_list") {
					return fmt.Errorf("銀行、カード詳細ページに遷移できませんでした: %s", illegalCheck)
				}

				chromedp.Nodes(`#js-btn-acts-more`, &continuationNode, chromedp.AtLeast(0)).Do(ctx)
				for len(continuationNode) != 0 {
					chromedp.Click(`#js-btn-acts-more`, chromedp.NodeVisible).Do(ctx)
					chromedp.WaitVisible(`#js-acts-table-tbody`, chromedp.NodeVisible).Do(ctx)
					time.Sleep(1 * time.Second)
					chromedp.Nodes(`#js-btn-acts-more`, &continuationNode, chromedp.AtLeast(0)).Do(ctx)
				}
				chromedp.Nodes(`.ca-general-container`, &bankIdNode, chromedp.ByQuery).Do(ctx)
				res, err = dom.GetOuterHTML().WithNodeID(bankIdNode[0].NodeID).Do(ctx)
				if err != nil {
					fmt.Printf("error %s", err)
				}
				bankId, err := scrapingOfDetailBankId(res)
				if err != nil {
					fmt.Printf("error %s", err)
				}
				chromedp.Nodes(`.ca-table`, &detailsNode, chromedp.ByQueryAll).Do(ctx)
				if len(detailsNode) == 0 {
					return fmt.Errorf("銀行、並びにカード情報が取れませんでした。")
				}
				res, err := dom.GetOuterHTML().WithNodeID(detailsNode[0].NodeID).Do(ctx)
				if err != nil {
					fmt.Printf("error %s", err)
				}

				details, err = scrapingOfDetails(res, details, bankId)
				if err != nil {
					fmt.Printf("error %s", err)
				}

			}
			return nil
		})

		chromedp.Run(ctx, fetchBanks, fetchDetails)

		return nil

	})

	// crawling開始
	err := chromedp.Run(ctx,
		loginActionFunc,
		crawlingActionFunc,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return user, banks, details, err
}

func scrapingOfOfficeName(res string, user []*User) ([]*User, error) {
	readerCurContents := strings.NewReader(res)
	contentsDom, err := goquery.NewDocumentFromReader(readerCurContents)
	if err != nil {
		return nil, err
	}
	user = append(user, &User{officeName: contentsDom.Find("#dropdown-office").Text()})
	return user, nil
}

func scrapingOfBanks(res string, banks []*Bank, user []*User, detailUrls []*DetailUrl) ([]*Bank, []*DetailUrl, error) {
	readerCurContents := strings.NewReader(res)
	contentsDom, err := goquery.NewDocumentFromReader(readerCurContents)

	if err != nil {
		return nil, nil, err
	}
	contentsDom.Find(`tr.js-account-row`).Each(func(i int, v *goquery.Selection) {
		bankName := v.Find("td").Text()
		strA := v.Find("td.text-right").Text()
		strA = strings.Replace(strA, "円", "", -1)
		strA = strings.Replace(strA, ",", "", -1)
		amount, _ := strconv.ParseInt(strA, 10, 64)

		bankStatus := v.Find("td:nth-child(5) span span").Text()
		if bankStatus != "正常" {
			bankStatus = "取得中"
		}

		bankId, _ := v.Find("td:nth-child(6)").Attr("id")

		kindURL, _ := v.Find("td:nth-child(8) a").Attr("href")
		kindURL = "https://accounting.moneyforward.com" + kindURL

		banks = append(banks, &Bank{
			Id:              uuid.NewString(),
			BankId:          strings.Replace(bankId, "js-edit-account-", "", 1),
			BankName:        bankName[:strings.Index(bankName, "（")],
			Amount:          amount,
			OfficeName:      user[0].officeName,
			LastCommit:      v.Find("td:nth-child(3)").Text(),
			ResitrationDate: v.Find("td.last-aggregated-datetime.text-center").Text(),
			BankStatus:      bankStatus,
		})
		detailUrls = append(detailUrls, &DetailUrl{
			DetailUrl: kindURL,
		})
	})
	return banks, detailUrls, nil
}

func fetchBanksKind(ctx context.Context, detailUrl string, registeredListUrl string) (string, error) {
	var illegalCheck string
	var kind string
	bankKindNodes := []*cdp.Node{}

	chromedp.Navigate(detailUrl).Do(ctx)
	chromedp.WaitVisible(`#js-acts-table-tbody`, chromedp.NodeVisible).Do(ctx)
	chromedp.Location(&illegalCheck).Do(ctx)
	if !strings.Contains(illegalCheck, "https://accounting.moneyforward.com/accounts/trans_list") {
		return "", fmt.Errorf("銀行、カード詳細ページに遷移できませんでした: %s", illegalCheck)
	}

	chromedp.Nodes(`.sub-account-txt`, &bankKindNodes, chromedp.ByQueryAll).Do(ctx)
	if len(bankKindNodes) == 0 {
		return "", fmt.Errorf("銀行、並びにカード情報のNode取れませんでした。")
	}
	res, err := dom.GetOuterHTML().WithNodeID(bankKindNodes[0].NodeID).Do(ctx)
	if err != nil {
		return "", fmt.Errorf("銀行、並びにカード情報のdomが取れませんでした。 %s", err)
	}
	chromedp.Text(`.sub-account-txt`, &kind, chromedp.NodeVisible).Do(ctx)
	chromedp.Navigate(registeredListUrl).Do(ctx)
	chromedp.Location(&illegalCheck).Do(ctx)
	if !strings.Contains(illegalCheck, "https://accounting.moneyforward.com/accounts") {
		return "", fmt.Errorf("明細の登録済み一覧ページに遷移できませんでした: %s", illegalCheck)
	}
	chromedp.WaitVisible(`body`, chromedp.ByQuery).Do(ctx)
	kind, err = scrapingOfBankKind(res)
	if err != nil {
		return "", fmt.Errorf("銀行、並びにカード情報が取れませんでした。 %s", err)
	}

	return kind, nil
}

func scrapingOfBankKind(res string) (string, error) {
	readerCurContents := strings.NewReader(res)
	contentsDom, err := goquery.NewDocumentFromReader(readerCurContents)
	if err != nil {
		return "", err
	}
	kind := contentsDom.Find(".sub-account-txt").Text()
	if strings.Contains(kind, "支店") {
		kind = "銀行"
	} else {
		kind = "クレジットカード"
	}
	return kind, nil
}

func scrapingOfDetails(res string, details []*Detail, bankId string) ([]*Detail, error) {
	readerCurContents := strings.NewReader(res)
	contentsDom, err := goquery.NewDocumentFromReader(readerCurContents)

	if err != nil {
		return nil, err
	}
	contentsDom.Find(`#js-acts-table-tbody > tr`).Each(func(i int, v *goquery.Selection) {

		strP := v.Find("td:nth-child(4)").Text()
		strP = strings.Replace(strP, "円", "", -1)
		strP = strings.Replace(strP, ",", "", -1)
		payment, _ := strconv.ParseInt(strP, 10, 64)

		strA := v.Find("td:nth-child(5)").Text()
		strA = strings.Replace(strA, "円", "", -1)
		strA = strings.Replace(strA, ",", "", -1)
		amount, _ := strconv.ParseInt(strA, 10, 64)

		details = append(details, &Detail{
			BankId:            bankId,
			TradingDate:       v.Find("td:nth-child(2)").Text(),
			TradingContent:    v.Find("td:nth-child(3)").Text(),
			Payment:           payment,
			Amount:            amount,
			BankName:          v.Find("td:nth-child(6)").Text(),
			Status:            v.Find("td:nth-child(7)").Text(),
			TransactionNumber: v.Find("td:nth-child(8)").Text(),
			edit:              v.Find("td:nth-child(9)").Text(),
		})
	})
	return details, nil
}

func scrapingOfDetailBankId(res string) (string, error) {
	readerCurContents := strings.NewReader(res)
	contentsDom, err := goquery.NewDocumentFromReader(readerCurContents)
	if err != nil {
		return "", err
	}
	bankId, _ := contentsDom.Find(".js-account-searchable-input").Attr("data-accountid")
	return bankId, nil
}
