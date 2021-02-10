package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"testCDN/model"
)

func (h *Handler) SortFile(c *gin.Context) {
	//data := struct {
	//	FilePath string `json:"file_path"`
	//}{}
	//if err := c.BindJSON(&data); err != nil {
	//	h.errorEndCall(c, http.StatusBadRequest, "FormatError", "Sort: BindJson: %v", err)
	//	return
	//}
	//var path string
	//if data.FilePath == "" {
	//	path = "config/config.json"
	//} else {
	//	path = data.FilePath
	//}
	path := "config/config.json"
	file, err := os.Open(path)
	if err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "Cannot open the file", "Sort: Open: %v", err)
		return
	}
	defer file.Close()
	config := model.Config{}
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "FormatError", "Sort: BindJson: %v", err)
		return
	}
	id, err := h.controller.SortSubnet(config)
	if err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "InternalError", "Sort: SortSubnet: %v", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) Sort(c *gin.Context) {
	config := model.Config{}
	if err := c.BindJSON(&config); err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "FormatError", "Sort: BindJson: %v", err)
		return
	}
	id, err := h.controller.SortSubnet(config)
	if err != nil {
		h.errorEndCall(c, http.StatusBadRequest, "InternalError", "Sort: SortSubnet: %v", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}
