package factorymethod;

import java.io.BufferedReader;
import java.io.FileReader;
import java.util.Iterator;

public class PDFTextExtractor implements ITextExtractor {

    public PDFTextExtractor(String path) {
        //BufferedReader reader = new BufferedReader(new FileReader(path));
    }

    @Override
    public String ReadLine() {
        return null;
    }

    @Override
    public Iterator<String> ReadLines() {
        return null;
    }

}