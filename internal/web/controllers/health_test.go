package controllers

import (
	"testing"
)

func TestGetHealth(t *testing.T) {
	/*
		resp := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(resp)
		GetHealth(c)
		assert.Equal(t, 200, resp.Code)

		var got gin.H
		err := json.Unmarshal(resp.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}

		want := gin.H{"status": "ok"}

		assert.Equal(t, want, got)
	*/
}
