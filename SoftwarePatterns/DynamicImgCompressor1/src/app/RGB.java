package app;

import interfaces.IPixelData;

public class RGB implements IPixelData {
    
    public RGB(int red, int green, int blue) {
        Red = red;
        Green = green;
        Blue = blue;
    }

    public int Red, Green, Blue;
}