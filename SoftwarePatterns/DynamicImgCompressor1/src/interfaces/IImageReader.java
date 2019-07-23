package interfaces;

import java.util.Iterator;

public interface IImageReader
{
    Iterator<IPixel> getPixels();
    IPixel getAverage(IPixel NE, IPixel SE, IPixel SW, IPixel NW);
    
}