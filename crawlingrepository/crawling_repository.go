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
	"github.com/gyoza-and-beer/proto-MF/config"
	"github.com/gyoza-and-beer/proto-MF/crawlingproto"
)

type CrawlingInterface interface {
	Crawling(pass string, input *crawlingproto.UserInput) ([]*User, []*Bank, []*Detail, error)
}

type crawlingRepository struct{}

func NewCrawling() CrawlingInterface {
	return &crawlingRepository{}
}

type User struct {
	Id        string
	UserId    string
	LastId    string
	UpdatedAt time.Time
}

type Bank struct {
	Id              string
	UserId          string
	BankId          string
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
	BankName          string
	Status            string
	UpdatedDate       string
	GettingDate       string
	TransactionNumber string
	Crawling          time.Time
}

// スクレイピング時に必要な事業所名、銀行口座名、銀行口座ID
// var bankNameAndId []map[string]string
// var officeName string

func (c *crawlingRepository) Crawling(pass string, input *crawlingproto.UserInput) ([]*User, []*Bank, []*Detail, error) {
	var user []*User
	var banks []*Bank
	var details []*Detail

	ctx, cancel := context.WithTimeout(config.NewChromedpContext(), 3*time.Minute)
	defer cancel()

	loginURL := "https://erp.moneyforward.com/session/new"
	topURL := "https://erp.moneyforward.com/home"
	illegalCheck := ""
	// illegalCheckNode := []*cdp.Node{}

	loginIdSel := `/html/body/main/div/div/div/div/div[2]/div/div[2]/div[1]/section/form/div[2]/div/input`
	loginPassSel := `/html/body/main/div/div/div/div/div[2]/div/div[2]/div[1]/section/form/div[2]/div/input[2]`
	nextLoginButtonSel := `/html/body/main/div/div/div/div/div[2]/div/div[2]/div[1]/section/form/div[2]/div/div[3]/input`

	// topLoginButtonSel := `/html/body/main/div/div/div/div/div[2]/div/div[2]/div[1]/section/form/div[2]/div/div[3]/input`

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

			time.Sleep(1 * time.Second)

			chromedp.Nodes(`.ca-table-horizontal.mf-mb10`, &bankNodes, chromedp.ByQueryAll).Do(ctx)
			if len(bankNodes) == 0 {
				return fmt.Errorf("銀行、並びにカード情報が取れませんでした。")
			}
			res, err := dom.GetOuterHTML().WithNodeID(bankNodes[0].NodeID).Do(ctx)
			if err != nil {
				fmt.Printf("error %s", err)
			}

			banks, err = scrapingOfBanks(res, banks)
			if err != nil {
				fmt.Printf("error %s", err)
			}
			for i, n := range banks {
				var detailURL = n.Kind
				kind, err := fetchBanksKind(ctx, detailURL, registeredListUrl)
				if err != nil {
					return fmt.Errorf("銀行、並びにカードの種別ができませんでした。")
				}
				banks[i].Kind = kind
			}
			return nil
		})
		chromedp.Run(ctx, fetchBanks)

		// banks, err := fetchBanks(ctx, banks, registeredListUrl)
		// for _, bank := range banks {
		// 	fmt.Println(bank.Kind)
		// }
		// fmt.Println(err)
		// err := chromedp.Run(ctx, getBanksActionFunc, getDetailActionFunc)
		// if err != nil {
		// 	return err
		// }

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

func scrapingOfBanks(res string, banks []*Bank) ([]*Bank, error) {
	readerCurContents := strings.NewReader(res)
	contentsDom, err := goquery.NewDocumentFromReader(readerCurContents)

	if err != nil {
		return nil, err
	}
	contentsDom.Find(`tr.js-account-row`).Each(func(i int, v *goquery.Selection) {
		bankName := v.Find("td").Text()

		strAmount := v.Find("td.text-right").Text()
		strAmount = strings.Replace(strAmount, "円", "", -1)
		strAmount = strings.Replace(strAmount, ",", "", -1)
		amount, _ := strconv.ParseInt(strAmount, 10, 64)

		bankStatus := v.Find("td:nth-child(5) span span").Text()
		if bankStatus != "正常" {
			bankStatus = "取得中"
		}

		kindURL, _ := v.Find("td:nth-child(8) a").Attr("href")
		kindURL = "https://accounting.moneyforward.com" + kindURL

		banks = append(banks, &Bank{
			Id:              uuid.NewString(),
			BankName:        bankName[:strings.Index(bankName, "（")],
			Amount:          amount,
			LastCommit:      v.Find("td:nth-child(3)").Text(),
			ResitrationDate: v.Find("td.last-aggregated-datetime.text-center").Text(),
			BankStatus:      bankStatus,
			Kind:            kindURL,
		})
	})
	return banks, nil
}

func fetchBanksKind(ctx context.Context, detailUrl string, registeredListUrl string) (string, error) {
	var illegalCheck string
	var kind string

	chromedp.Navigate(detailUrl).Do(ctx)
	chromedp.WaitVisible(`#js-acts-table-tbody`, chromedp.NodeVisible).Do(ctx)
	chromedp.Location(&illegalCheck).Do(ctx)
	if !strings.Contains(illegalCheck, "https://accounting.moneyforward.com/accounts/trans_list") {
		return "", fmt.Errorf("銀行、カード詳細ページに遷移できませんでした: %s", illegalCheck)
	}
	chromedp.Text(`.sub-account-txt`, &kind, chromedp.NodeVisible).Do(ctx)
	chromedp.Navigate(registeredListUrl).Do(ctx)
	chromedp.Location(&illegalCheck).Do(ctx)
	if !strings.Contains(illegalCheck, "https://accounting.moneyforward.com/accounts") {
		return "", fmt.Errorf("明細の登録済み一覧ページに遷移できませんでした: %s", illegalCheck)
	}
	chromedp.WaitVisible(`body`, chromedp.ByQuery).Do(ctx)
	if strings.Contains(kind, "支店") {
		kind = "銀行"
	} else {
		kind = "クレジットカード"
	}
	return kind, nil
}
