
interface Person{
    public void eat();
    public void sleep();
}

class Student implements Person{

    @Override
    public void eat() {
        System.out.println("学生是人，也要吃饭");
    }

    @Override
    public void sleep() {
        System.out.println("学生是人，也要睡觉");
    }
}
public class Interface {
    public static void main(String[] args){
        Student stu = new Student();
        stu.eat();
        stu.sleep();
    }
}
