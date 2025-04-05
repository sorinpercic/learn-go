🔧 How a Web Scraper Works

    Send a Request
    The scraper sends an HTTP request to a webpage (like your browser does when you visit a site).

    Get the Response (HTML)
    The server returns the HTML of the page. This is the raw content of the webpage.

    Parse the HTML
    The scraper uses a parser (like BeautifulSoup in Python) to understand the structure of the HTML.

    Extract Data
    It finds the elements it needs (e.g., product names, prices, titles) using HTML tags or CSS selectors.

    Store the Data
    The extracted data is saved in a structured format—CSV, JSON, or a database.