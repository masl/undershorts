import { Button } from "@/components/ui/button"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"

export default function ShortenForm() {
  const formSchema = z.object({
    longUrl: z.string().url({
      message: "Invalid URL",
    }),
  })

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      longUrl: "",
    },
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    // TODO: send to endpoint
    console.log(values)
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="longUrl"
          render={({ field }) => (
            <FormItem>
              <FormLabel>URL</FormLabel>
              <FormControl>
                <Input placeholder="http://github.com/masl/undershorts" {...field} />
              </FormControl>
              <FormDescription>This is your URL that will be shortened.</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Shorten</Button>
      </form>
    </Form>
  )
}
