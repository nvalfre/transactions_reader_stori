package email_service

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"transactions_reader_stori/domain"
)

func (s *EmailService) buildEmailBodyContent(summary *domain.SummaryVO) string {
	sort.Slice(summary.TransactionSummary, func(i, j int) bool {
		return summary.TransactionSummary[i].Month.Before(summary.TransactionSummary[j].Month)
	})

	var sb strings.Builder
	sbPointer := &sb

	sbPointer.WriteString("<html>")
	s.buildHeader(sbPointer)
	s.buildBody(summary, sbPointer)
	sbPointer.WriteString("</html>")

	body := sbPointer.String()
	return body
}

func (s *EmailService) buildBody(summary *domain.SummaryVO, sb *strings.Builder) {
	sb.WriteString("<body>")
	s.buildBodyContainer(summary, sb)
	sb.WriteString("</body>")
}

func (s *EmailService) buildBodyContainer(summary *domain.SummaryVO, sb *strings.Builder) {
	sb.WriteString(`<div class="container">`)
	s.buildAppbarHeader(sb)

	s.buildTotalBalanceSection(summary, sb)

	s.buildTransactionSummarySection(summary, sb)

	s.buildAverageCreditAmountSection(summary, sb)

	s.buildAverageDebitAmountSection(summary, sb)

	s.buildMonthlySummary(summary, sb)

	sb.WriteString(`</div>`)
}

func (s *EmailService) buildMonthlySummary(summary *domain.SummaryVO, sb *strings.Builder) {
	for month, monthlySummary := range summary.MonthlySummary {
		sb.WriteString(`<div class="section">`)
		sb.WriteString(fmt.Sprintf(`<h2>Monthly Summary for %s</h2>`, month.String()))
		sb.WriteString(`<ul class="section-content">`)
		for _, ts := range monthlySummary.Summaries {
			sb.WriteString(`<li>`)
			sb.WriteString(fmt.Sprintf(`Transaction ID: %d, Amount: %.2f`, ts.ID, ts.Amount))
			sb.WriteString(`</li>`)
		}
		sb.WriteString(`</ul>`)
		sb.WriteString(`</div>`)
	}
}

func (s *EmailService) buildHeader(sb *strings.Builder) {
	sb.WriteString("<head><style>")
	sb.WriteString(style)
	sb.WriteString("</style></head>")
}

func (s *EmailService) buildAppbarHeader(sb *strings.Builder) {
	sb.WriteString(`<h1>Summary Report</h1>`)
}

func (s *EmailService) buildTotalBalanceSection(summary *domain.SummaryVO, sb *strings.Builder) {
	sb.WriteString(`<div class="section">`)
	sb.WriteString(`<h2>Total Balance</h2>`)
	sb.WriteString(fmt.Sprintf(`<p>%.2f</p>`, summary.TotalBalance))
	sb.WriteString(`</div>`)
}

func (s *EmailService) buildTransactionSummarySection(summary *domain.SummaryVO, sb *strings.Builder) {
	sb.WriteString(`<div class="section">`)
	sb.WriteString(`<h2>Transaction Summaries</h2>`)
	prevMonth := time.Time{}
	for _, ts := range summary.TransactionSummary {
		if !ts.Month.Equal(prevMonth) {
			sb.WriteString(fmt.Sprintf(`<h3>%s</h3>`, ts.Month.Format("January")))
			prevMonth = ts.Month
		}
		sb.WriteString(`<ul><li>`)
		sb.WriteString(fmt.Sprintf(`Transaction ID: %d, Amount: %.2f`, ts.ID, ts.Amount))
		sb.WriteString(`</li></ul>`)
	}
	sb.WriteString(`</div>`)
}

func (s *EmailService) buildAverageDebitAmountSection(summary *domain.SummaryVO, sb *strings.Builder) {
	sb.WriteString(`<div class="section">`)
	sb.WriteString(`<h2>Average Debit Amount</h2>`)
	sb.WriteString(fmt.Sprintf(`<p>%.2f</p>`, summary.AverageDebit))
	sb.WriteString(`</div>`)
}

func (s *EmailService) buildAverageCreditAmountSection(summary *domain.SummaryVO, sb *strings.Builder) {
	sb.WriteString(`<div class="section">`)
	sb.WriteString(`<h2>Average Credit Amount</h2>`)
	sb.WriteString(fmt.Sprintf(`<p>%.2f</p>`, summary.AverageCredit))
	sb.WriteString(`</div>`)
}
