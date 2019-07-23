package app;

import java.io.File;
import java.util.Iterator;

import interfaces.IPixel;

public class App {
    public static void main(String[] args) throws Exception {
        File file = new File("Z:\\9 Sem\\Software Patterns\\DynamicImgCompressor1\\resources\\sample.jpg");
        JPEGReader jr = new JPEGReader(file);

        int count = 0;
        Iterator<IPixel> pi = jr.getPixels();
        while (pi.hasNext()) {
            pi.next();
            count++;
        }
        System.out.print(count);
    }
}