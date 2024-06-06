package reports

import (
	"fmt"
	"log"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/jung-kurt/gofpdf"
)

func GenerateDailyReport(data []dto.ReportDTO) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	tr := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.Cell(40, 10, tr("Relatório Diário de Vendas"))
	pdf.SetFont("Arial", "", 12)

	go func() {
		for _, sale := range data {
			line := fmt.Sprintf("Venda ID: %d, Produto ID: %d, Cliente ID: %d, Quantidade: %d, Data: %s",
				sale.Id, sale.ProductId, sale.ClientId, sale.Quantity, sale.Date.Format("01-01-2006"))

			pdf.Ln(10)
			pdf.Cell(40, 10, line)
		}
	}()

	err := pdf.OutputFileAndClose("relatorio_diario.pdf")
	if err != nil {
		log.Println("Error generating PDF:", err)
	} else {
		log.Println("Relatório diário gerado com sucesso!")
	}
}
