class Person {
  String firstName;
  String lastName;
  int age;
  Person(this.firstName, this.lastName, this.age);

  String get fullName => '$firstName $lastName';
}

void main() {
  var person = Person('John', 'Doe', 30);
  print(person.fullName);
}
