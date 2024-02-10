"use client";

import { useCreatePaymentMethod } from "@/modules/payment/useCreatePaymentMethod";
import {
  Button,
  chakra,
  Flex,
  FormControl,
  FormLabel,
  Input,
} from "@chakra-ui/react";
import { GMOMultiPaymentScript } from "./script";

export const CreatePaymentMethodForm = () => {
  const { form, onCreatePaymentMethod } = useCreatePaymentMethod();

  return (
    <chakra.form
      borderWidth={1}
      maxW={600}
      mx={"auto"}
      px={4}
      py={5}
      onSubmit={onCreatePaymentMethod}
    >
      <GMOMultiPaymentScript
        apiKey={process.env.NEXT_PUBLIC_GMO_TOKEN_API_KEY!}
      />
      <FormControl>
        <FormLabel>カード番号</FormLabel>
        <Input
          type={"number"}
          {...form.register("cardno", { valueAsNumber: true })}
        />
      </FormControl>
      <Flex gap={2}>
        <FormControl>
          <FormLabel>年</FormLabel>
          <Input
            type={"number"}
            {...form.register("expireYear", { valueAsNumber: true })}
          />
        </FormControl>
        <FormControl>
          <FormLabel>月</FormLabel>
          <Input
            type={"number"}
            {...form.register("expireMonth", { valueAsNumber: true })}
          />
        </FormControl>
      </Flex>
      <FormControl>
        <FormLabel>名義人</FormLabel>
        <Input {...form.register("holdername")} />
      </FormControl>
      <FormControl>
        <FormLabel>セキュリティコード</FormLabel>
        <Input
          maxW={"6em"}
          textAlign={"right"}
          type={"number"}
          {...form.register("securitycode", { valueAsNumber: true })}
        />
      </FormControl>
      <Flex justify={"center"} mt={3}>
        <Button
          colorScheme={"teal"}
          minW={"10em"}
          type={"submit"}
          variant={"outline"}
        >
          送信
        </Button>
      </Flex>
    </chakra.form>
  );
};
