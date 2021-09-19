import React from 'react'
import {createStackNavigator} from '@react-navigation/stack'
import {Detail, Pokemons } from '../pages';
const Stack = createStackNavigator();


const Router = () =>{
    return(
        <Stack.Navigator>
            <Stack.Screen 
                name="Pokemons"
                component={Pokemons}
                options={{headerShown:false}}
            />
            <Stack.Screen 
                name="Details"
                component={Detail}
                options={{headerShown:false}}
            />
            
        </Stack.Navigator>
    )
}

export default Router