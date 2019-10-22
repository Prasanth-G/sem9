import java.util.Dictionary;
import java.util.Hashtable;

class APIRequest {

    public String Request;
    public int AccessToken;

    public APIRequest(String request, int accessToken) {
        Request = request;
        AccessToken = accessToken;
    }
}

class ResourceServer {
    AuthorizationServer as;

    ResourceServer(AuthorizationServer as) {
        this.as = as;
    }

    public String RequestHandler(APIRequest request) {
        if (as.isAccessTokenValid(request.AccessToken)) {
            return "Request Accepted";
        } else {
            return "Invalid AccessToken";
        }
    }
}

class AuthorizationServer {

    private Dictionary userData;

    AuthorizationServer() {
        userData = new Hashtable();
    }

    public int GetAccessToken(String username, String password) {
        int h = (username + password).hashCode();
        userData.put(h, username);
        return h;
    }

    public boolean isAccessTokenValid(int token) {
        return userData.get(token) != null;
    }
}

class Client {
    private ResourceServer rs;
    int token;

    Client (ResourceServer rs, AuthorizationServer as, String username, String password) {
        this.rs = rs;
        token = as.GetAccessToken(username, password);
    }

    public String SendRequest(String request) {
        return rs.RequestHandler(new APIRequest(request, token));
    }

}

public class App2
{
    public static void main(String[] args)
    {
        AuthorizationServer as = new AuthorizationServer();
        AuthorizationServer as1 = new AuthorizationServer();
        ResourceServer rs = new ResourceServer(as);

        Client c1 = new Client(rs, as, "Client1", "pass1");
        Client c2 = new Client(rs, as1, "Client2", "pass2");

        System.out.println(c1.SendRequest("From Client1"));
        System.out.println(c2.SendRequest("From Client1"));
    }
}