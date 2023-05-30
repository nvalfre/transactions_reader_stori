package email_service

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"transactions_reader_stori/domain"
)

const (
	sectionTitle      = "Total Balance"
	sectionColor      = "#3366CC"
	contentColor      = "#333333"
	tableHeaderColor  = "#CCCCCC"
	tableContentColor = "#FFFFFF"
	textColor         = "#000000"
)

func (s *EmailService) buildEmailBodyContent(summary *domain.SummaryVO) string {
	sort.Slice(summary.TransactionSummary, func(i, j int) bool {
		return summary.TransactionSummary[i].Month.Before(summary.TransactionSummary[j].Month)
	})

	var sb strings.Builder
	sbPointer := &sb

	sbPointer.WriteString("<html>")
	sbPointer.WriteString("<head>")
	sbPointer.WriteString(`<meta name="viewport" content="width=device-width, initial-scale=1.0">`)
	sbPointer.WriteString("</head>")
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

	s.buildTotalBalanceSection(summary, sb, sectionTitle, sectionColor, contentColor)
	s.buildTransactionSummarySection(summary, sb, sectionTitle, sectionColor, contentColor)
	s.buildAverageDebitAmountSection(summary, sb, sectionTitle, sectionColor, contentColor)
	s.buildAverageCreditAmountSection(summary, sb, sectionTitle, sectionColor, contentColor)
	s.buildMonthlySummary(summary, sb, sectionColor, tableHeaderColor, tableContentColor, textColor)

	sb.WriteString(`</div>`)
}

func (s *EmailService) buildAppbarHeader(sb *strings.Builder) {
	sb.WriteString(`<h1 style="color: #333333;">Summary Report</h1>`)
}

func (s *EmailService) buildTotalBalanceSection(summary *domain.SummaryVO, sb *strings.Builder, sectionTitle, sectionColor, contentColor string) {
	sb.WriteString(`<div class="section">`)
	sb.WriteString(fmt.Sprintf(`<h2 style="color: %s;">%s</h2>`, sectionColor, sectionTitle))
	sb.WriteString(fmt.Sprintf(`<p style="color: %s;">%.2f</p>`, contentColor, summary.TotalBalance))
	sb.WriteString(`</div>`)
}

func (s *EmailService) buildTransactionSummarySection(summary *domain.SummaryVO, sb *strings.Builder, sectionTitle, sectionColor, contentColor string) {
	sb.WriteString(`<div class="section">`)
	sb.WriteString(fmt.Sprintf(`<h2 style="color: %s;">%s</h2>`, sectionColor, sectionTitle))
	prevMonth := time.Time{}
	for _, ts := range summary.TransactionSummary {
		if !ts.Month.Equal(prevMonth) {
			sb.WriteString(fmt.Sprintf(`<h3 style="color: %s;">%s</h3>`, contentColor, ts.Month.Format("January")))
			prevMonth = ts.Month
		}
		sb.WriteString(`<ul><li>`)
		sb.WriteString(fmt.Sprintf(`<span style="color: %s;">Transaction ID: %d, Amount: %.2f</span>`, contentColor, ts.ID, ts.Amount))
		sb.WriteString(`</li></ul>`)
	}
	sb.WriteString(`</div>`)
}

func (s *EmailService) buildAverageDebitAmountSection(summary *domain.SummaryVO, sb *strings.Builder, sectionTitle, sectionColor, contentColor string) {
	sb.WriteString(`<div class="section">`)
	sb.WriteString(fmt.Sprintf(`<h2 style="color: %s;">%s</h2>`, sectionColor, sectionTitle))
	sb.WriteString(fmt.Sprintf(`<p style="color: %s;">%.2f</p>`, contentColor, summary.AverageDebit))
	sb.WriteString(`</div>`)
}

func (s *EmailService) buildAverageCreditAmountSection(summary *domain.SummaryVO, sb *strings.Builder, sectionTitle, sectionColor, contentColor string) {
	sb.WriteString(`<div class="section">`)
	sb.WriteString(fmt.Sprintf(`<h2 style="color: %s;">%s</h2>`, sectionColor, sectionTitle))
	sb.WriteString(fmt.Sprintf(`<p style="color: %s;">%.2f</p>`, contentColor, summary.AverageCredit))
	sb.WriteString(`</div>`)
}

func (s *EmailService) buildMonthlySummary(summary *domain.SummaryVO, sb *strings.Builder, sectionColor, tableHeaderColor, tableContentColor, textColor string) {
	for month, monthlySummary := range summary.MonthlySummary {
		sb.WriteString(`<div class="section">`)
		sb.WriteString(fmt.Sprintf(`<h2>Monthly Summary for %s</h2>`, month.String()))
		sb.WriteString(`<table class="summary-table">`)
		sb.WriteString(`<tr>`)
		sb.WriteString(fmt.Sprintf(`<th style="background-color: %s; color: %s;">Month</th>`, tableHeaderColor, textColor))
		sb.WriteString(fmt.Sprintf(`<th style="background-color: %s; color: %s;">Transaction ID</th>`, tableHeaderColor, textColor))
		sb.WriteString(fmt.Sprintf(`<th style="background-color: %s; color: %s;">Amount</th>`, tableHeaderColor, textColor))
		sb.WriteString(`</tr>`)
		for _, ts := range monthlySummary.Summaries {
			sb.WriteString(`<tr>`)
			sb.WriteString(fmt.Sprintf(`<td style="background-color: %s; color: %s;">%s</td>`, tableContentColor, textColor, month.String()))
			sb.WriteString(fmt.Sprintf(`<td style="background-color: %s; color: %s;">%d</td>`, tableContentColor, textColor, ts.ID))
			sb.WriteString(fmt.Sprintf(`<td style="background-color: %s; color: %s;">%.2f</td>`, tableContentColor, textColor, ts.Amount))
			sb.WriteString(`</tr>`)
		}
		sb.WriteString(`</table>`)
		sb.WriteString(`</div>`)
	}
}
