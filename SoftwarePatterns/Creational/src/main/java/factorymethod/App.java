package factorymethod;

import java.io.IOException;

import org.apache.tika.exception.TikaException;
import org.xml.sax.SAXException;

/**
 * Hello world!
 */
public final class App {
    private App() {
    }

    /**
     * Says hello to the world.
     * 
     * @param args The arguments of the program.
     * @throws TikaException
     * @throws SAXException
     * @throws IOException
     */
    public static void main(String[] args) throws IOException, SAXException, TikaException {
        String path = "C:\\Users\\TEMP.CS2K16.012\\Downloads\\sample.html";
		ITextExtractor file = TextExtractorFactory.getReader(FileReaders.WebpageReader, path);
        System.out.println(file.ReadLine());
        System.out.println("END");
    }
}
