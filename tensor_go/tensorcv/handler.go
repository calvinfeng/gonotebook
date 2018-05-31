package tensorcv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Response defines the structure of a HTTP JSON response to client.
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// NewImageRecognitionHandler returns a HTTP handler that will handle a request to perform image
// recognition.
func NewImageRecognitionHandler(labels map[int]string, modelPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		imgFile, header, err := r.FormFile("image")
		if err != nil {
			response := &Response{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}

			if resBytes, err := json.Marshal(response); err == nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write(resBytes)
			}
			return
		}

		imgName := strings.Split(header.Filename, ".")

		var imgBuffer bytes.Buffer
		io.Copy(&imgBuffer, imgFile)
		fmt.Printf("Received image %s which as %d bytes\n", imgName, len(imgBuffer.Bytes()))

		var imgFormat string
		if imgName[1] == "jpeg" || imgName[1] == "jpg" {
			imgFormat = "jpeg"
		} else {
			imgFormat = "png"
		}

		imgTensor, err := GetTensorFromImageBuffer(imgBuffer, imgFormat, 3)
		fmt.Println("Image tensor is loaded:", imgTensor.Shape())

		softmaxScore := RunResNetModel(imgTensor, modelPath)
		if softmaxScore != nil {
			selected := make(map[int]bool)
			classIdx, prob := ArgMax(softmaxScore[0], selected)
			fmt.Printf("Predicted class is %s with %.2f probability\n", labels[classIdx], prob)

			response := &Response{
				Status: http.StatusOK,
				Message: fmt.Sprintf(
					"Predicted class is %s with %.2f probability",
					labels[classIdx], prob),
			}

			if resBytes, err := json.Marshal(response); err == nil {
				w.WriteHeader(http.StatusOK)
				w.Write(resBytes)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// ArgMax takes in a list and return the index and value of the max element.
func ArgMax(list []float32, selected map[int]bool) (int, float32) {
	idx := 0
	max := list[idx]
	for i, el := range list {
		if el > max && !selected[i] {
			idx = i
			max = el
		}
	}

	selected[idx] = true

	return idx, max
}
