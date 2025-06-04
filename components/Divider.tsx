import { ShortcutProps, defaultShortcuts } from '@/styles/shortcuts';
import React from 'react';
import { View } from 'react-native';

export interface DividerProps extends ShortcutProps { }

export const Divider = (props: DividerProps) => {
  return (
    <View style={ [defaultShortcuts(props), {
      backgroundColor: 'lightgray',
      height: 1,
    }] } />
  );
};