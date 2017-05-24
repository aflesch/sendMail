# sendMail
Example go application to build HTML message and send it by AWS SES

## Template
This tool uses a HTML template with parameter "Name".
It rather easy to add more parameters - should be defined in the template.html file, and also in Params struct.

## Mail
To send mail, we are using AWS SES service.
AWS SES is currently supported only in the following regions:
* eu-west-1 (Ireland)
* us-east-1 (N. Virginia)
* us-west-2 (Oregon)

There are three methods to send email using AWS SES:
1. Amazon SES console
2. SMTP Interface
3. Amazon SES API

We are using the third method - SES API: You need to create AWS session and then create SES session.
 
You need to register to SES in your AWS account.

Then, you need to verify your domain name "https://eu-west-1.console.aws.amazon.com/ses/home?region=eu-west-1#verified-senders-domain:" - you can use all mail addresses in the domain.  
Or, you can verify e-mail address 'https://eu-west-1.console.aws.amazon.com/ses/home?region=eu-west-1#verified-senders-email:' - you can use only this mail.  

#### Notice
In the beginning, Amazon will put you in "SES sandbox" - that's mean:
* You are limited in the number of e-mails per day (200)
* You can send one email per second
* <b>Recpients also must be verified</b> - you need to open a case to Amazon to let you out from the Sandbox

q