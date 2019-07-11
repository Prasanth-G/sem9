package app;

import java.io.IOException;
import org.apache.lucene.queryparser.classic.QueryParser;
import org.apache.lucene.analysis.standard.StandardAnalyzer;

public class App {
    public static void main(String[] args) throws Exception {
        System.out.println("Hello Java");

        QueryParser qp = new QueryParser("firstName", new StandardAnalyzer());
        System.out.println(qp);
    }

    // private static IndexWriter createWriter() throws IOException
    // {
    //     FSDirectory dir = FSDirectory.open(Paths.get(INDEX_DIR));
    //     IndexWriterConfig config = new IndexWriterConfig(new StandardAnalyzer());
    //     IndexWriter writer = new IndexWriter(dir, config);
    //     return writer;
    // }
}