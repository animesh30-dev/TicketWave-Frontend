import { ShortcutProps, defaultShortcuts } from "@/styles/shortcuts";
import React, { PropsWithChildren } from 'react';
import { Text as RNText, TextProps } from 'react-native';

export interface CustomTextProps extends PropsWithChildren, ShortcutProps, TextProps {
  fontSize?: number;
  bold?: boolean;
  underline?: boolean;
  color?: string;
}

export const Text = ({ fontSize = 18, bold, underline, color, children, ...restProps }: CustomTextProps) => {
  return (
    <RNText
      style={ [defaultShortcuts(restProps), {
        fontSize,
        fontWeight: bold ? 'bold' : 'normal',
        textDecorationLine: underline ? 'underline' : 'none',
        color,
      }] }
      { ...restProps }
    >
      { children }
    </RNText>
  );
};