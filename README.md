# SellarAppTask

Install mysql server with userid = 'root' and password = 'root'.

Run the main file with command: "go run main.go" 

Use Postman to hit request in below format:

http://localhost:10000/scrap_url

{
	"url":"https://www.amazon.in/SFT-Raisins-Munakka-Superior-Quality/dp/B075X2QSQD/?_encoding=UTF8&pd_rd_w=4vTzK&pf_rd_p=91a3a6e7-aee3-435e-9a78-c619cecf54a5&pf_rd_r=3173DZJQ9MZ8T7TMZ51A&pd_rd_r=962d4154-3b2d-4586-8dfe-15ac66758dc4&pd_rd_wg=SqKRX&ref_=pd_gw_trq_rep_sims_gw"
}

Run the following command in mysql command line, to see the inserted data.

Use scrapped;

Select * from scrappeddata;
