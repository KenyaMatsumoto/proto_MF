package crawlingrepository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/gyoza-and-beer/proto-MF/crawlingproto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CrawlingReadInterface interface {
	UserRead(ctx context.Context, req *pb.MfRequest) (users []*pb.Office, err error)
	BankRead(ctx context.Context, req *pb.MfRequest, officeName string, startDay string, lastDay string) (resultBanks []*pb.Bank, err error)
	CardRead(ctx context.Context, req *pb.MfRequest, officeName string, startDay string, lastDay string) (resultCards []*pb.Card, err error)
	detailRead(ctx context.Context, userId string, bankId string, officeName string, startDay string, lastDay string) (details []*pb.Detail, err error)
}

type CrawlingRead struct {
	client *sql.DB
}

type Date struct {
	detailDate time.Time
	updatedAt  time.Time
}

func NewCrawlingRead(db *sql.DB) CrawlingReadInterface {
	return &CrawlingRead{db}
}

func (c *CrawlingRead) UserRead(ctx context.Context, req *pb.MfRequest) (offices []*pb.Office, err error) {
	row, err := c.client.Query(OfficeReadStmt(), req.UserInput.UserId)
	if err != nil {
		return nil, fmt.Errorf("事業所名のクエリ取得に失敗しました: %s", err)
	}
	defer row.Close()

	for row.Next() {
		user := &pb.Office{}
		date := &Date{}
		err := row.Scan(&user.OfficeName, &date.updatedAt)
		if err != nil {
			return nil, fmt.Errorf("事業所名の取得に失敗しました: %s", err)
		}

		user.Crawling = timestamppb.New(date.updatedAt)
		offices = append(offices, user)
	}

	return offices, nil
}

func (c *CrawlingRead) BankRead(ctx context.Context, req *pb.MfRequest, officeName string, startDay string, lastDay string) (resultBanks []*pb.Bank, err error) {
	var banks []*pb.Bank
	resultBanks, err = getBankDetails(c, ctx, req.UserInput.UserId, officeName, startDay, lastDay, banks)
	if err != nil {
		return nil, err
	}
	return resultBanks, nil
}

func (c *CrawlingRead) CardRead(ctx context.Context, req *pb.MfRequest, officeName string, startDay string, lastDay string) (resultCards []*pb.Card, err error) {

	var cards []*pb.Card
	resultCards, err = getCardDetails(c, ctx, req.UserInput.UserId, officeName, startDay, lastDay, cards)
	if err != nil {
		return nil, err
	}
	return resultCards, nil
}

func getBankDetails(c *CrawlingRead, ctx context.Context, userId string, officeName string, startDay string, lastDay string, banks []*pb.Bank) ([]*pb.Bank, error) {

	rows, err := c.client.Query(DistinctBankIdAndBankNameStmt(), userId, officeName)
	if err != nil {
		return nil, fmt.Errorf("銀行IDと銀行名の取得クエリの作成に失敗しました: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		bank := &pb.Bank{}
		if err := rows.Scan(&bank.BankId, &bank.BankName, &bank.BankStatus, &bank.BankLastCommit); err != nil {
			return nil, fmt.Errorf("%sの明細取得に失敗しました！: %s", bank.BankName, err)
		}
		details, err := c.detailRead(ctx, userId, bank.BankId, officeName, startDay, lastDay)
		if err != nil {
			return nil, fmt.Errorf("%sの明細取得に失敗しました!!: %s", bank.BankName, err)
		}
		bank.Details = details
		banks = append(banks, bank)
	}
	return banks, nil
}

func (c *CrawlingRead) detailRead(ctx context.Context, userId string, bankId string, officeName string, startDay string, lastDay string) (details []*pb.Detail, err error) {
	rows, err := c.client.Query(DetailStmt(startDay, lastDay), userId, bankId)
	if err != nil {
		return nil, fmt.Errorf("明細のクエリ取得に失敗しました: %s", err)
	}
	defer rows.Close()

	details = []*pb.Detail{}
	for rows.Next() {
		detail := &pb.Detail{}
		date := &Date{}

		if err := rows.Scan(&detail.DetailName, &detail.Contents, &detail.Amount, &detail.Balance, &date.detailDate); err != nil {
			return nil, fmt.Errorf("明細の取得に失敗しました。: %s", err)
		}
		detail.DetailDate = timestamppb.New(date.detailDate)

		details = append(details, detail)
	}
	return details, nil
}

func getCardDetails(c *CrawlingRead, ctx context.Context, userId string, officeName string, startDay string, lastDay string, cards []*pb.Card) ([]*pb.Card, error) {
	rows, err := c.client.Query(DistinctCardIdAndCardNameStmt(), userId, officeName)
	if err != nil {
		return nil, fmt.Errorf("各クレジットカードと各クレジットカード残高クエリ取得に失敗しました: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		card := &pb.Card{}

		if err := rows.Scan(&card.CardId, &card.CardName, &card.CardStatus, &card.CardLastCommit); err != nil {
			return nil, fmt.Errorf("%sの明細取得に失敗しました: %s", card.CardName, err)
		}
		details, err := c.detailRead(ctx, userId, card.CardId, officeName, startDay, lastDay)
		if err != nil {
			return nil, fmt.Errorf("%sの明細取得に失敗しました: %s", card.CardName, err)
		}
		card.Details = details
		cards = append(cards, card)
	}
	return cards, nil
}
