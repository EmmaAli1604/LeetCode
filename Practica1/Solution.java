package Practica1;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

class Solution {
    /*
     * @param un arreglo de números enteros
     * @return un arreglo de números enteros
     * Este método lo que hace es que dado un arreglo numérico multiplica cada 
     * uno de sus elementos excepto por el í-esimo elemento, y el reultado lo 
     * guardo en un arreglo de números enteros
     */
    public static int[] productExceptSelf(int[] nums) {
        int i; //Conatdor de los for
        int aux = 1; //Variable auxiliar 
        int aux2=1; //Variable auxiliar 2
        
        /*Lista para cuando recorremos de izquierda a derecha por primera vez el 
        arreglo y multiplicamos en dos en dos los elementos de este.
        Se utilizan ya que podemos acceder al i-esimo elemento con el método get() */
        List<Integer> list = new ArrayList<Integer>(); 
        /*Segunda lista para cuando recorremos de derecha a izquierda por segunda 
        vez el arrgelo multiplicando en dos en dos los elementos de este*/
        List<Integer> list2 = new ArrayList<Integer>();

        /* Arreglo que se va a regresar con la respuesta*/
        int[] answer = new int[nums.length];
        /*En este for recorremos de izquierda a derecha el arreglo dado en dos en dos 
        y el reultado se va guardando en la lista */
        for(i = 0 ; i < nums.length ; i++){
            aux = aux * nums[i];
            list.add(aux);
        }

        /*En este for recorremos de derecha a izquierda el arreglo dado en dos en 
        dos y el reultado se va guardando en la lista*/
        for(i = nums.length -1 ; i >= 0; i--){
            aux2 = aux2 * nums[i];
            list2.add(aux2);
        }

        /*Volteamos la lista2 para poder multiplicar bien los elementos */
        Collections.reverse(list2);

        /*Recorremos el arreglo del parametro para poder identificar en que lugar 
        estan los elementos y ver por cual cantidad se debe de multiplicar el el 
        i-esimo elemento menos 1 en la fila con el i-esimo elemento más 1 de la 
        segunda lista  */
        for(i = 0; i < nums.length; i++){
            if(i == 0){
                answer[i]=(int) list2.get(i+1);
            }
            else if( i > 0  && i < nums.length -1){
                answer[i]=list.get(i-1) * list2.get(i+1);
            }
            else{
                answer[i]=(int)list.get(i-1);
            }
        }

        return answer;
    }

}