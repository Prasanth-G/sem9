package jar;

/**
 * Hello world!
 *
 */

import edu.uci.ics.crawler4j.crawler.*;
import edu.uci.ics.crawler4j.fetcher.PageFetcher;
import edu.uci.ics.crawler4j.robotstxt.RobotstxtConfig;
import edu.uci.ics.crawler4j.robotstxt.RobotstxtServer;
import org.apache.lucene.queryparser.classic.QueryParser;
import org.apache.lucene.analysis.standard.StandardAnalyzer;

public class App 
{
    public static void main( String[] args ) throws Exception
    {
        //QueryParser qp = new QueryParser("firstName", new StandardAnalyzer());
        //System.out.println(qp);
        
        String crawlStorageFolder = "/data/crawl/root";
        //String crawlStorageFolder = "D:\\data\\crawl\\root";
        int numberOfCrawlers = 10, maxPage = 10, MAX_CRAWL_DEPTH = 1;

        CrawlConfig config = new CrawlConfig();
        config.setCrawlStorageFolder(crawlStorageFolder);
        config.setMaxPagesToFetch(maxPage);
        config.setMaxDepthOfCrawling(MAX_CRAWL_DEPTH);

        // Instantiate the controller for this crawl.
        PageFetcher pageFetcher = new PageFetcher(config);
        RobotstxtConfig robotstxtConfig = new RobotstxtConfig();
        RobotstxtServer robotstxtServer = new RobotstxtServer(robotstxtConfig, pageFetcher);
        CrawlController controller = new CrawlController(config, pageFetcher, robotstxtServer);

        // For each crawl, you need to add some seed urls. These are the first
        // URLs that are fetched and then the crawler starts following links
        // which are found in these pages
        controller.addSeed("https://en.wikipedia.org/wiki/Computer_science");
    	
    	// The factory which creates instances of crawlers.
        CrawlController.WebCrawlerFactory<IRCrawler> factory = () -> {
            return new IRCrawler();
        };
        
        //Wait until crawlling completed
        controller.start(factory, numberOfCrawlers);
    }
}
