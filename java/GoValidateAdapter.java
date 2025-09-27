public class GoValidateAdapter implements Validate {

    static {
        System.loadLibrary("govalidate");
    }

    private native boolean Validate(String input);

    @Override
    public boolean validate(String input) {
        return Validate(input);
    }
}
