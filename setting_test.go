package telegraph_test

//func TestSetWebHookSuccess(t *testing.T) {
//	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
//		"ok": true,
//		"result": true,
//		"description": "Webhook was set"
//	}`)
//	defer gock.Off()
//
//	client := telegraph.NewClient("token")
//
//	res, body, err := client.SetWebHook("https://telegram.squarecode.co.id/v1/webhook").Commit()
//	assert.NotNil(t, res)
//	assert.Equal(t, http.StatusOK, res.StatusCode)
//	assert.NotNil(t, body)
//	assert.NoError(t, err)
//}
//
//func TestSetWebHookWithCertificate(t *testing.T) {
//	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
//		"ok": true,
//		"result": true,
//		"description": "Webhook was set"
//	}`)
//	defer gock.Off()
//
//	client := telegraph.NewClient("token")
//
//	res, body, err := client.SetWebHook("https://telegram.squarecode.co.id/v1/webhook").
//		Certificate("./README.md").Commit()
//	assert.NotNil(t, res)
//	assert.Equal(t, http.StatusOK, res.StatusCode)
//	assert.NotNil(t, body)
//	assert.NoError(t, err)
//}
//
//func TestSetWebHookWithMaxConnection(t *testing.T) {
//	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
//		"ok": true,
//		"result": true,
//		"description": "Webhook was set"
//	}`)
//	defer gock.Off()
//
//	client := telegraph.NewClient("token")
//
//	res, body, err := client.SetWebHook("https://telegram.squarecode.co.id/v1/webhook").
//		MaxConnection(100).Commit()
//	assert.NotNil(t, res)
//	assert.Equal(t, http.StatusOK, res.StatusCode)
//	assert.NotNil(t, body)
//	assert.NoError(t, err)
//}
//
//func TestSetWebHookWithAllowedUpdates(t *testing.T) {
//	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
//		"ok": true,
//		"result": true,
//		"description": "Webhook was set"
//	}`)
//	defer gock.Off()
//
//	client := telegraph.NewClient("token")
//
//	res, body, err := client.SetWebHook("https://telegram.squarecode.co.id/v1/webhook").
//		AllowedUpdate("update", "notify", "message").Commit()
//	assert.NotNil(t, res)
//	assert.Equal(t, http.StatusOK, res.StatusCode)
//	assert.NotNil(t, body)
//	assert.NoError(t, err)
//}
//
//func TestSetWebHookWithAllOptions(t *testing.T) {
//	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSetWebHook, "token")).Reply(http.StatusOK).JSON(`{
//		"ok": true,
//		"result": true,
//		"description": "Webhook was set"
//	}`)
//	defer gock.Off()
//
//	client := telegraph.NewClient("token")
//
//	res, body, err := client.SetWebHook("https://telegram.squarecode.co.id/v1/webhook").
//		AllowedUpdate("update", "notify", "message").MaxConnection(100).Certificate("./LICENSE").
//		Commit()
//	assert.NotNil(t, res)
//	assert.Equal(t, http.StatusOK, res.StatusCode)
//	assert.NotNil(t, body)
//	assert.NoError(t, err)
//}
