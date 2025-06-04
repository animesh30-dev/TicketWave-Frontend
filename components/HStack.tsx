import { Stack, StackProps } from "@/components/Stack";
import React from "react";

interface HStackProps extends StackProps { }

export function HStack(props: HStackProps) {
  return (
    <Stack { ...props } direction="row">
      { props.children }
    </Stack>
  );
}