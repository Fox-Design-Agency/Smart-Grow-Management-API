// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growingmediumhandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"totally-legit-grow-management/v1/internal/helpers"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

// EditGrowingMedium will handle the initial request and then send the formed request
// to the logic layer before returning the request
func (s *GrowingMedium) EditGrowingMedium(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.EditGrowingMediumRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}
	/**********************************************************************
	/
	/	Call Logic Layer
	/
	/**********************************************************************/
	resp, err := s.Logic.EditGrowingMedium(&form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}
	/**********************************************************************
	/
	/	Return Success
	/
	/**********************************************************************/
	helpers.SendSuccessHeader(w, resp)
}
