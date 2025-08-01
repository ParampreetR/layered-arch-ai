package master_models

import "time"

// PraoMst represents the MASTER.PRAO_MST table
// PRAO stands for Pr/Pay and Accounts Office (state/regional level)
// Verified by Parampreet
type PraoMst struct {
	// Primary key - DDL uses PRAO_ID_PK INTEGER NOT NULL
	PraoIDPk int `gorm:"column:prao_id_pk;primaryKey;autoIncrement" json:"prao_id_pk"`

	// Core identification fields - exact DDL column specifications
	PraoAIN  int    `gorm:"column:prao_ain;not null;uniqueIndex" json:"prao_ain"`       // PRAO_AIN INTEGER NOT NULL
	PraoCode string `gorm:"column:prao_code;type:varchar(6);not null" json:"prao_code"` // PRAO_CODE VARCHAR(6) NOT NULL
	PraoType string `gorm:"column:prao_type;type:char(2);not null" json:"prao_type"`    // PRAO_TYPE CHAR(2) NOT NULL
	PraoName string `gorm:"column:prao_name;type:varchar(75)" json:"prao_name"`         // PRAO_NAME VARCHAR(75)

	// Government and sector classification
	PraoGovtType    string `gorm:"column:prao_govt_type;type:char(3)" json:"prao_govt_type"`     // PRAO_GOVT_TYPE CHAR(3)
	PraoModelCode   string `gorm:"column:prao_model_code;type:char(2)" json:"prao_model_code"`   // PRAO_MODEL_CODE CHAR(2)
	PraoSectorCateg int    `gorm:"column:prao_sector_categ" json:"prao_sector_categ"`            // PRAO_SECTOR_CATEG INTEGER
	PraoRegnCode    string `gorm:"column:prao_regn_code;type:varchar(10)" json:"prao_regn_code"` // PRAO_REGN_CODE VARCHAR(10)

	// Office information
	PraoOffc     string `gorm:"column:prao_offc;type:varchar(75)" json:"prao_offc"`           // PRAO_OFFC VARCHAR(75)
	PraoOfcrName string `gorm:"column:prao_ofcr_name;type:varchar(75)" json:"prao_ofcr_name"` // PRAO_OFCR_NAME VARCHAR(75)

	// Address details
	PraoAddrLn1 string  `gorm:"column:prao_addr_ln1;type:varchar(30)" json:"prao_addr_ln1"` // PRAO_ADDR_LN1 VARCHAR(30)
	PraoAddrLn2 *string `gorm:"column:prao_addr_ln2;type:varchar(30)" json:"prao_addr_ln2"` // PRAO_ADDR_LN2 VARCHAR(30)
	PraoAddrLn3 *string `gorm:"column:prao_addr_ln3;type:varchar(30)" json:"prao_addr_ln3"` // PRAO_ADDR_LN3 VARCHAR(30)
	PraoAddrLn4 string  `gorm:"column:prao_addr_ln4;type:varchar(30)" json:"prao_addr_ln4"` // PRAO_ADDR_LN4 VARCHAR(30)
	PraoStateCd string  `gorm:"column:prao_state_cd;type:char(2)" json:"prao_state_cd"`     // PRAO_STATE_CD CHAR(2)
	PraoCounCd  string  `gorm:"column:prao_coun_cd;type:char(2)" json:"prao_coun_cd"`       // PRAO_COUN_CD CHAR(2)
	PraoPinCd   int     `gorm:"column:prao_pin_cd" json:"prao_pin_cd"`                      // PRAO_PIN_CD INTEGER

	// Contact information
	PraoTelNum    string `gorm:"column:prao_tel_num;type:varchar(13)" json:"prao_tel_num"`         // PRAO_TEL_NUM VARCHAR(13)
	PraoAltTelNum string `gorm:"column:prao_alt_tel_num;type:varchar(13)" json:"prao_alt_tel_num"` // PRAO_ALT_TEL_NUM VARCHAR(13)
	PraoFaxNum    string `gorm:"column:prao_fax_num;type:varchar(13)" json:"prao_fax_num"`         // PRAO_FAX_NUM VARCHAR(13)
	PraoEmail     string `gorm:"column:prao_email;type:varchar(80)" json:"prao_email"`             // PRAO_EMAIL VARCHAR(80)

	// Status and activation
	PraoActiveFlg string     `gorm:"column:prao_active_flg;type:char(1)" json:"prao_active_flg"` // PRAO_ACTIVE_FLG CHAR(1)
	PraoActvnDt   *time.Time `gorm:"column:prao_actvn_dt;type:date" json:"prao_actvn_dt"`        // PRAO_ACTVN_DT DATE
	PraoDeactvnDt *time.Time `gorm:"column:prao_deactvn_dt;type:date" json:"prao_deactvn_dt"`    // PRAO_DEACTVN_DT DATE
	PraoSts       int        `gorm:"column:prao_sts;not null" json:"prao_sts"`                   // PRAO_STS INTEGER NOT NULL

	// Audit fields - DDL patterns
	PraoCrtdBy     string     `gorm:"column:prao_crtd_by;type:varchar(12);not null" json:"prao_crtd_by"`       // PRAO_CRTD_BY VARCHAR(12) NOT NULL
	PraoCrtdTmstmp time.Time  `gorm:"column:prao_crtd_tmstmp;type:timestamp;not null" json:"prao_crtd_tmstmp"` // PRAO_CRTD_TMSTMP TIMESTAMP NOT NULL
	PraoMdfdBy     *string    `gorm:"column:prao_mdfd_by;type:varchar(12)" json:"prao_mdfd_by"`                // PRAO_MDFD_BY VARCHAR(12)
	PraoMdfdTmstmp *time.Time `gorm:"column:prao_mdfd_tmstmp;type:timestamp" json:"prao_mdfd_tmstmp"`          // PRAO_MDFD_TMSTMP TIMESTAMP
	PraoAckNo      int        `gorm:"column:prao_ack_no;not null" json:"prao_ack_no"`                          // PRAO_ACK_NO INTEGER NOT NULL
}

// TableName sets the table name for GORM
func (PraoMst) TableName() string {
	return `"master"."prao_mst"`
}
