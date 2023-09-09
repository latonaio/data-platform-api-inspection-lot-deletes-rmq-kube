package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-inspection-lot-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-inspection-lot-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) HeaderRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.InspectionLot = %d ", input.Header.InspectionLot)
	rows, err := c.db.Query(
		`SELECT 
			header.InspectionLot
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_header_data as header 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) InspectionsRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Inspection {
	where := fmt.Sprintf("WHERE inspection.InspectionLot IS NOT NULL\nAND header.InspectionLot = %d", input.Header.InspectionLot)
	rows, err := c.db.Query(
		`SELECT 
			inspection.InspectionLot, inspection.Inspection
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_inspection_data as inspection
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_inspection_lot_header_data as header
		ON header.InspectionLot = inspection.InspectionLot ` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToInspection(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
