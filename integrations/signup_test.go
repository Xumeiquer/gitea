// Copyright 2017 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/setting"

	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	assert.NoError(t, models.LoadFixtures())
	setting.Service.EnableCaptcha = false

	req, err := http.NewRequest("POST", "/user/sign_up",
		bytes.NewBufferString(url.Values{
			"user_name": []string{"exampleUser"},
			"email":     []string{"exampleUser@example.com"},
			"password":  []string{"examplePassword"},
			"retype":    []string{"examplePassword"},
		}.Encode()),
	)
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp := MakeRequest(req)
	assert.EqualValues(t, http.StatusFound, resp.HeaderCode)

	// should be able to view new user's page
	req, err = http.NewRequest("GET", "/exampleUser", nil)
	assert.NoError(t, err)
	resp = MakeRequest(req)
	assert.EqualValues(t, http.StatusOK, resp.HeaderCode)
}
