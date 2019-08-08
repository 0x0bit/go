
class Animal {
    public   String name;
    public String subject;

    void eat(String food) {
        System.out.println(name + "喜欢吃：" + food + ",它属于：" + subject);
    }
}

class Cat extends Animal{
    public int age;
    void sleep() {
        System.out.println(name + "今年" + age + "岁了，特别喜欢睡觉");
    }
}

public class Test{
    public static void main(String[] args){
        Animal a = new Animal();
        a.name = "动物";
        a.subject = "动物科";
        a.eat("肉");

        Cat cat = new Cat();
        cat.name = "咪咪";
        cat.subject = "猫科";
        cat.age = 1;
        cat.eat("鱼");
        cat.sleep();
    }
}