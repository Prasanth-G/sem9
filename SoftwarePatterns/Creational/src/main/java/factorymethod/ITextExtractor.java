package factorymethod;

import java.util.Iterator;

// Need     : If user wants to create an instance from a set of subclasses which may override the parent class implementation 
//              and the choice depends on state of running application, config settings.
// Disadvantages : High coupling of other components to service provider

public interface ITextExtractor {
    String ReadLine();
    Iterator<String> ReadLines();
}