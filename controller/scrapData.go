package controller
import (
    "fmt"
    "log"
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "strings"
    "regexp"
    "SELLARAPP/model"
)

var data model.Scrapped_data

func Get_data(scrap_url string) (model.Scrapped_data){

    client := &http.Client{}
    req, _ := http.NewRequest("GET", scrap_url, nil)
    // header is required since website is popping robot check dialog box.
    req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.90 Safari/537.36")
    response, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer response.Body.Close()
    
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }
    // Find all links and process them with the function
    // defined earlier
    document.Find("div").Each(processElement)
    document.Find("span").Each(getTitle)
    return data
}

func getTitle(index int, element *goquery.Selection){

    spanId, exists := element.Attr("id")
    if exists{
        if spanId == "productTitle"{
            data.Title = element.Text()
            data.Title = strings.Join(strings.Fields(data.Title), " ")
            fmt.Println(data.Title)
        }
        if spanId == "acrCustomerReviewText"{
            data.Total_reviews = element.Text()
            data.Total_reviews = strings.Join(strings.Fields(data.Total_reviews), " ")
            fmt.Println(data.Total_reviews)
        }
    }

}

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
    // See if the href attribute exists on the element
    div_id, exists := element.Attr("id")
    if exists {
        if div_id == "imgTagWrapperId"{ 
            data.Image_url = element.Find("img").AttrOr(`data-a-dynamic-image`, ``)
            r, _ := regexp.Compile(`:\[\d+,\d+\]`)
            var images = r.ReplaceAllString(data.Image_url, "")
            data.Image_url = images
            fmt.Println(images)
        }
        if (div_id == "olp_feature_div" || div_id == "price"){
            data.Price = strings.TrimSpace(element.Text())
            data.Price = strings.Join(strings.Fields(data.Price), " ")
            if div_id == "price"{
                r := regexp.MustCompile(`Price: (.) \d*(\.)\d*`)
                data.Price = r.FindString(data.Price)
            }
            if div_id == "olp_feature_div"{
                r, _ := regexp.Compile(`(.) \d*(\.)\d*`)
                data.Price = r.FindString(data.Price)
            }
            fmt.Println(data.Price)
        }
        if (div_id == "feature-bullets"){
            data.Description = strings.TrimSpace(element.Text())
            data.Description = strings.Join(strings.Fields(data.Description), " ")
            fmt.Println(data.Description)
        }
    }
}