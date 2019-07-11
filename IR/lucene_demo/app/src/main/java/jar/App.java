package jar;

/**
 * Hello world!
 *
 */

import java.io.IOException;
import org.apache.lucene.queryparser.classic.QueryParser;
import org.apache.lucene.analysis.standard.StandardAnalyzer;

public class App 
{
    public static void main( String[] args ) throws Exception
    {
        System.out.println( "Hello World!" );
        QueryParser qp = new QueryParser("firstName", new StandardAnalyzer());
        System.out.println(qp);
    }
}
