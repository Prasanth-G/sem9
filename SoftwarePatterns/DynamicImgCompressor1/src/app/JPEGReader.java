package app;

import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import java.net.URL;
import java.util.Iterator;

import javax.imageio.*;
import interfaces.*;

public class JPEGReader implements IImageReader {

    private BufferedImage image;

    public JPEGReader(File file) throws IOException {
        this.image = ImageIO.read(file);
    }

    public JPEGReader(URL url) throws IOException {
        this.image = ImageIO.read(url);
    }

    @Override
    public Iterator<IPixel> getPixels() {
        return new RGBPixelIterator(image);
    }

    @Override
    public IPixel getAverage(IPixel NE, IPixel SE, IPixel SW, IPixel NW) {
        return null;
    }

}