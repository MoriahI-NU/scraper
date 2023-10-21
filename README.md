# Web Scraping With Colly

## Problem Setup:
A tech firm is creating its own online library. Management will collect information for their project by scraping Wikipedia pages. The current program uses Python and Scrapy, but it runs very slowly (example code for Python/Scrapy was analyzed at NWSPS). Our job is to create a program in Go that will provide the same results with a faster runtime.

The methods I chose to go with use either a Colly framework or a GoWiki framework to scrape information with golang. In this repo you'll find a couple of different attempts. Outputs (in json lines format) and .exe files are included for each one.

1. Run Main.exe
This attempt produces a result that matches closest with the Python/Scrapy output. Each article object includes the URL, Title, Text, and Category Tags. They are also listed on a separate lines in the jsonl file. However, the text output includes attributes relating to text format (like italics, alignment, etc.) which seems unnecessary for the scope of this project. It is included here because it shows that golang can get essentially the same results as Python in less time. On my system the Main.exe ran in approximately three counts, while the Python script took about eight counts.

Output of main.exe = output3.jsonl

2. Run clean_alt.exe
This was created to produce a cleaner output than the first program. Here the objects include the article title with distinct subheaders:section content. There is no text formatting information which will make it easier to read and gather the information that actually matters. However, this attempt is not perfect as it queries for headers and paragraphs. Since elements are not nested within the sections that contain them, any information contained in tables/bullet lists do not result in the output file. I am working on getting these things to be included, but for now this .exe is able to relay most of the pertinent information.

Output of clean_alt.exe = output.jsonl

3. Run gowiki.exe
This last attempt is probably the cleanest of all. It is the only one that does not use colly, but instead uses the gowiki package which is tailored for scraping wikipedia pages. The output from this one includes all information (from paragraphs, table information, bullet list information, etc.) and excludes any formatting/superfluous text. This attempt is the best one for scraping article contents and returning *clean* and *comprehensive* text. The code is also more concise than that which uses colly.

Output of gowiki.exe = output2.jsonl