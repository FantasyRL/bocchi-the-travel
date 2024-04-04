package sender

import (
	"bocchi/config"
	"bocchi/rpc/user/dal/db"
	"bytes"
	"encoding/base64"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

func SendEmail(user *db.User, buf *bytes.Buffer) error {
	qrcode := make([]byte, 114514)                 //开辟存储空间
	base64.StdEncoding.Encode(qrcode, buf.Bytes()) //buff转成base64

	e := &email.Email{
		To:      []string{user.Email},
		From:    config.Sender.From,
		Subject: "您 Steam 愿望单上的 《天使☆纷扰 RE-BOOT!》无法购买！",
		HTML:    BuildSendMessage(user.UserName, qrcode),
		Headers: textproto.MIMEHeader{},
	}

	//因为使用base64内嵌所以不需要附件了
	//fileName := fmt.Sprintf("%s_%d.png", user.Email, user.ID)
	//attachment, err := e.Attach(buf, fileName, "image/png; charset=utf-8")
	//if err != nil {
	//	return err
	//}
	//attachment.HTMLRelated = true

	return e.Send(config.Sender.Host+":"+config.Sender.Port, smtp.PlainAuth("", config.Sender.Sender, config.Sender.Password, config.Sender.Host))
}
func BuildSendMessage(name string, base64png []byte) []byte {
	msg1 := `
	<title>Title</title>
	<table width="100%" border="0" cellspacing="0" cellpadding="0">
	   <tbody><tr>
	       <td class="p-80 mpy-35 mpx-15" bgcolor="#212429" style="padding: 80px;">
	           <table width="100%" border="0" cellspacing="0" cellpadding="0">
	
	
	               <tbody><tr>
	                   <td class="img pb-45" style="font-size:0pt; line-height:0pt; text-align:left; padding-bottom: 45px;">
	                       <a href="https://www.yuanshen.com/" target="_blank" rel="noopener">
	                           <img src="https://www.yuanshen.com/images/ys.96a55539.png" width="112" height="88" border="0">
	                       </a>
	
	                   </td>
	               </tr>
	
	
	
	               <tr>
	                   <td>
	
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="title-36 pb-30 c-grey6 fw-b" style="font-size:36px; line-height:42px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-bottom: 30px; color:#bfbfbf; font-weight:bold;"><span style="color: #77b9ee;">`
	msg2 := `</span></td>
	                           </tr>
	                           </tbody></table>
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="text-18 c-grey4 pb-30" style="font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; color:#dbdbdb; padding-bottom: 30px;">看起来您正在尝试登录。此处是您访问帐户所需的OTP：</td>
	                           </tr>
	                           </tbody></table>
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="pb-70 mpb-50" style="padding-bottom: 70px;">
	                                   <table width="100%" border="0" cellspacing="0" cellpadding="0" bgcolor="#17191c">
	                                       <tbody><tr>
	                                           <td class="py-30 px-56" style="padding-top: 30px; padding-bottom: 30px; padding-left: 56px; padding-right: 56px;">
	                                               <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                                                   <tbody><tr>
	                                                       <td style="font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; color:#8f98a0; text-align:center;">
	                                                           请求来自												</td>
	                                                   </tr>
	                                                   <tr>
	                                                       <td style="font-size:25px; line-height:30px; font-family:Arial, sans-serif, 'Motiva Sans'; color:#f1f1f1; text-align:center;letter-spacing:1px">
	                                                           提瓦特大陆												</td>
	                                                   </tr>
	                                                   <tr>
	                                                       <td style="padding-bottom: 16px"></td>
	                                                   </tr>
	                                                   </tbody><tbody><tr>
	                                                   <td class="img pb-45" style="font-size:0pt; line-height:0pt; text-align:center; padding-bottom: 45px;">
	                                                       <a>
	                                                           <img src="data:image/png;base64,`
	msg3 := `" width="200" height="200" align-items="center">
	                                                       </a>
	
	                                                   </td>
	                                               </tr>
	                                               </tbody></table>
	                                           </td>
	                                       </tr>
	                                       </tbody></table>
	                               </td>
	                           </tr>
	                           </tbody></table>
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="pb-30" style="padding-bottom: 30px;">
	                                   <table width="210" border="0" cellspacing="0" cellpadding="0">
	                                       <tbody><tr>
	                                           <td><br>&nbsp;</td>
	                                       </tr>
	                                       </tbody></table>
	                               </td>
	                           </tr>
	                           </tbody></table>
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="title-36 pb-30 c-grey6 fw-b" style="font-size:30px; line-height:34px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-bottom: 20px; color:#bfbfbf; font-weight:bold;">不是您？</td>
	                           </tr>
	                           </tbody></table>
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="text-18 c-grey4 pb-30" style="font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; color:#dbdbdb; padding-bottom: 30px;">您会收到这封电子邮件，是由于有人试图登录您的帐户，且提供了<span style="color: #ffffff; font-weight: bold;">正确的帐户名称与密码</span>。<br><br>
	                                   如果这不是您尝试登录，建议您<a href="https://www.yuanshen.com/" rel="noopener" target="_blank">进入此网站</a>。<br><br>
	                                   此电子邮件包含一个OTP，您需要用它访问您的帐户。切勿与任何人分享此OTP。</td>
	                           </tr>
	                           </tbody></table>
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="pb-30" style="padding-bottom: 30px;">
	                                   <table width="210" border="0" cellspacing="0" cellpadding="0">
	                                       <tbody><tr>
	                                           <td><br>&nbsp;</td>
	                                       </tr>
	                                       </tbody></table>
	                               </td>
	                           </tr>
	                           </tbody></table>
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="title-36 pb-30 c-grey6 fw-b" style="font-size:30px; line-height:34px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-bottom: 20px; color:#bfbfbf; font-weight:bold;">不住在提瓦特？</td>
	                           </tr>
	                           </tbody></table>
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="text-18 c-grey4 pb-30" style="font-size:18px; line-height:25px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; color:#dbdbdb; padding-bottom: 30px;">请先登录此网站<a href="https://www.yuanshen.com/" rel="noopener" target="_blank">(4天玩了82小时palworld后睡了327h)</a>再继续。</td>
	                           </tr>
	                           </tbody></table>
	
	
	                       <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                           <tbody><tr>
	                               <td class="pt-30" style="padding-top: 30px;">
	                                   <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                                       <tbody><tr>
	                                           <td class="img" width="3" bgcolor="#3a9aed" style="font-size:0pt; line-height:0pt; text-align:left;"></td>
	                                           <td class="img" width="37" style="font-size:0pt; line-height:0pt; text-align:left;"></td>
	                                           <td>
	                                               <table width="100%" border="0" cellspacing="0" cellpadding="0">
	                                                   <tbody><tr>
	                                                       <td class="text-16 py-20 c-grey4 fallback-font" style="font-size:16px; line-height:22px; font-family:Arial, sans-serif, 'Motiva Sans'; text-align:left; padding-top: 20px; padding-bottom: 20px; color:#f1f1f1;">
	                                                           祝您愉快，<br>
	                                                           以及你怎么知道我有5天早八                                                                                    </td>
	                                                   </tr>
	                                                   </tbody></table>
	                                           </td>
	                                       </tr>
	                                       </tbody></table>
	                               </td>
	                           </tr>
	                           </tbody></table>
	
	
	                   </td>
	               </tr>
	
	               </tbody></table>
	       </td>
	   </tr>
	
	   </tbody></table>`
	return bytes.Join([][]byte{[]byte(msg1), []byte(name), []byte(msg2), removeNullValue(base64png), []byte(msg3)}, []byte(""))
}

// removeNullValue 被空值气晕了
func removeNullValue(slice []byte) []byte {
	var output []byte
	for _, element := range slice {
		if element != '\000' {
			output = append(output, element)
		}
	}
	return output
}
