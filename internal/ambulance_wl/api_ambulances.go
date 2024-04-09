/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: a@a.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package ambulance_wl

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type AmbulancesAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // CreateAmbulance - Saves new ambulance definition
   CreateAmbulance(ctx *gin.Context)

    // DeleteAmbulance - Deletes specific ambulance
   DeleteAmbulance(ctx *gin.Context)

 }

 // partial implementation of AmbulancesAPI - all functions must be implemented in add on files
type implAmbulancesAPI struct {

}

func newAmbulancesAPI() AmbulancesAPI {
  return &implAmbulancesAPI{}
}

func (this *implAmbulancesAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/ambulance", this.CreateAmbulance)
  routerGroup.Handle( http.MethodDelete, "/ambulance/:ambulanceId", this.DeleteAmbulance)
}


// Copy following section to separate file, uncomment, and implement accordingly
