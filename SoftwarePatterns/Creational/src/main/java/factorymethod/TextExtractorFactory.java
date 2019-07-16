package factorymethod;

import java.io.IOException;

import org.apache.tika.exception.TikaException;
import org.xml.sax.SAXException;

public class TextExtractorFactory {
    public static ITextExtractor getReader(FileReaders type, String path)
            throws IOException, SAXException, TikaException {
        switch(type) {
            case SystemFileReader:
                return new PDFTextExtractor(path);
            case WebpageReader:
                return new HTMLTextExtractor(path);
            default:
                return null;
        }
    }
}