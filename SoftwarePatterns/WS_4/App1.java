interface CommonBC {
    public void Calculate();
}

class A {
    private CommonBC bc;

    A() {
        
    }

    public void doX(int x) {
        if (x < 10) {
            bc = new B();
            bc.Calculate();
        } else {
            bc = new C();
            bc.Calculate();
        }
    }
}

class B implements CommonBC {
    B() {
    }

    public void Calculate() {
        System.out.println("B Calculate");
    }
}

class C implements CommonBC {
    C() {
    }

    public void Calculate() {
        System.out.println("C Calculate");
    }
}

public class App1
{
    public static void main(String[] args)
    {
        A a = new A();
        a.doX(9);
    }
}