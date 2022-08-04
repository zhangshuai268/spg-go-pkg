package wechat

func (p *pay) PayAddCertFilePath(certFilePath string, keyFilePath string) error {
	p.CertFilePath = certFilePath
	p.KeyFilePath = keyFilePath
	return nil
}
