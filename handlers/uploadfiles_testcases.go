package handlers

import "net/url"

type UploaderTestCase struct {
	Name                 string
	ExpectedStatus       int
	ExpectedResponseBody string
	Form                 url.Values
}

var UploaderTestCases = []UploaderTestCase{
	{
		Name:                 "",
		ExpectedStatus:       200,
		ExpectedResponseBody: `{"file_name":"test.txt","source_path":"temp/test.txt","destination_path":"/","output_result":""}` + "\n",
		Form: url.Values{
			"host":             {"aspera.kaltura.com"},
			"password":         {"2qcttTOp"},
			"username":         {"asp2657331_1"},
			"auth_type":        {"password"},
			"source_pathS3URI": {"s3://aspera-delivery-temp/test.txt"},
			"destination_path": {`/`},
		},
	},
	{
		Name:                 "",
		ExpectedStatus:       200,
		ExpectedResponseBody: "Hello World!",
		Form: url.Values{
			"username":         {"xfer_sensical_amagi_trc"},
			"host":             {"aspera-hsts.sr.roku.com"},
			"port":             {"33001"},
			"auth_type":        {"privatekey"},
			"source_pathS3URI": {`s3://aspera-delivery-temp/FILE.xml`},
			"destination_path": {`/testing/D  `},
			"keyS3URI":         {`s3://dev-vod-processing/vod/Roku/keys/Roku_rsa`},
		},
	},
}
