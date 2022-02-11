public class collison_course {
  static int collision(List<Integer> speed, int pos) {
    int collisionPos = pos + speed.get(pos) - 1;
    int collisionCount = 0;

    for (int i = 0; i < speed.size(); i++) {
        if ((i + speed.get(i) - 1) >= collisionPos){
            collisionCount++;
        }
    }

    return collisionCount;
}
}
