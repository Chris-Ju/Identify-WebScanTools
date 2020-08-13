# Identify-WebScanTools

Go 语言实现识别 Web 扫描工具


```
AWVS

<1> Url:  
acunetix-wvs-test-for-some-inexistent-fileby_wvsacunetix_wvs_security_testacunetixacunetix_wvsacunetix_test
<2> Headers:
Acunetix-Aspect-Password:Cookie: acunetixCookieLocation: acunetix_wvs_security_testX-Forwarded-Host: acunetix_wvs_security_testX-Forwarded-For: acunetix_wvs_security_testHost: acunetix_wvs_security_testCookie: acunetix_wvs_security_testCookie: acunetix
Accept: acunetix/wvsOrigin: acunetix_wvs_security_testReferer: acunetix_wvs_security_testVia: acunetix_wvs_security_testAccept-Language: acunetix_wvs_security_testClient-IP: acunetix_wvs_security_testHTTP_AUTH_PASSWD: acunetixUser-Agent: acunetix_wvs_security_test
Acunetix-Aspect-Queries:任意值
Acunetix-Aspect:任意值
<3> Body:
acunetix_wvs_security_testacunetix
```

```
Netsparker

<1> URL:
netsparker
Netsparkerns: netsparker
<2> Headers:
X-Scanner: Netsparker
Location: Netsparker
Accept: netsparker/check
Cookie: netsparker
Cookie: NETSPARKER
<3> Body:
netsparker
```

```
Appscan

<1> URL:
UrlAppscan
<2> Headers:
Content-Type: Appscan
Content-Type: AppScan
HeaderAccept: Appscan
User-Agent:Appscan
<3> Body:
Appscan
```

```
Webinspect

<1> Url:
HP404
<2> Headers:
User-Agent: HP ASC
Cookie: webinspect
X-WIPP: 任意值
X-Request-Memo: 任意值
X-Scan-Memo: 任意值
Cookie: CustomCookie
X-RequestManager-Memo: 任意值
<3> Body:
Webinspect
```

```
Rsas (绿盟极光)

<1> Url:
nsfocus
<2> Headers:
User-Agent: Rsas
```

```
Nessus

<1> Url:
nessusNessus
<2> Headers:
x_forwarded_for: nessus
referer: nessus
host: nessus
<3> Body:
nessusNessus
```

```
WebReaver

<1> Headers:
User-Agent: WebReaver

```

```
Sqlmap

<1> Url:
sqlmap
<2> Headers:
User-Agent: sqlmap
<3> Body:
sqlmap
```