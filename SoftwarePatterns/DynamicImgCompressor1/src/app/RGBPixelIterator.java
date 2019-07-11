package app;

import java.awt.image.BufferedImage;
import java.util.Iterator;

import interfaces.IPixel;

public class RGBPixelIterator implements Iterator<IPixel> {

    private BufferedImage image;
    private int width, height, curr_x, curr_y;

    public RGBPixelIterator(BufferedImage image) {
        this.image = image;
        this.width = image.getWidth();
        this.height = image.getHeight();
        this.curr_x = this.curr_y = 0;
    }

    @Override
    public boolean hasNext() {
        if (curr_y < height) {
            return true;
        }
        return false;
    }

    @Override
    public IPixel next() {
        int  clr   = image.getRGB(curr_x, curr_y); 
        int  red   = (clr & 0x00ff0000) >> 16;
        int  green = (clr & 0x0000ff00) >> 8;
        int  blue  =  clr & 0x000000ff;
        IPixel output = new RGBPixel(curr_x, curr_y, new RGB(red, green, blue));
        if (curr_x < width) {
            curr_x++;
        } else {
            curr_x = 0;
            curr_y++;
        }
        return output;
    }
    
}