package alipay

// AgreementQuery https://docs.open.alipay.com/api_2/alipay.user.agreement.query
// 查询转账订单接口alipay.user.agreement.query
func (this *AliPay) AliPayAgreementQuery(param AliPayAgreementQuery) (results *AliPayAgreementQueryResponse, err error) {
	results = &AliPayAgreementQueryResponse{}
	err = this.doRequest("POST", param, results)
	return results, err
}

// AgreementQuery https://docs.open.alipay.com/api_2/alipay.user.agreement.query
// 查询转账订单接口alipay.user.agreement.query
func (this *AliPay) AliPayAgreementUnsign(param AliPayAgreementUnsign) (results *AliPayAgreementUnsignResponse, err error) {
	results = &AliPayAgreementUnsignResponse{}
	err = this.doRequest("POST", param, results)
	return results, err
}

// AgreementModify https://docs.open.alipay.com/api_2/alipay.user.agreement.executionplan.modify
// 修改协议接口alipay.user.agreement.executionplan.modify
func (this *AliPay) AliPayAgreementModify(param AliPayAgreementModify) (results *AliPayAgreementModifyResponse, err error) {
	results = &AliPayAgreementModifyResponse{}
	err = this.doRequest("POST", param, results)
	return results, err
}