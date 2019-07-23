package app;

import interfaces.IPixel;
import interfaces.IPixelData;

public class RGBPixel implements IPixel {

    private int x, y;
    private RGB rgbdata;

    public RGBPixel(int x, int y, RGB rgbdata) {
        this.x = x;
        this.y = y;
        this.rgbdata = rgbdata;
    }

    @Override
    public int getX() {
        return x;
    }

    @Override
    public int getY() {
        return y;
    }

    @Override
    public IPixelData getData() {
        return rgbdata;
    }
    
}