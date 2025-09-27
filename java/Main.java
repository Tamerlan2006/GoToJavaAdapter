public class Main {
    public static void main(String[] args) {
        Validate validator = new GoValidateAdapter();

        String test1 = "79927398713";  
        String test2 = "1234567890";    

        System.out.println("Test1: " + test1 + " → " + validator.validate(test1));
        System.out.println("Test2: " + test2 + " → " + validator.validate(test2));
    }
}
