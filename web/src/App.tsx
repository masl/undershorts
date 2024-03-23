import { Separator } from "@/components/ui/separator";
import ShortenForm from "@/components/ShortenForm";
import Header from "@/components/Header";

export default function App() {
  return (
    <>
      <div className="hidden space-y-6 p-10 pb-16 md:block">
        <Header />
        <Separator className="my-6" />
        <div className="flex justify-center">
          <ShortenForm />
        </div>
      </div>
    </>
  )

  // return (
  //   <>
  //     <h1 className="text-green-500">undershorts</h1>
  //     <div>
  //       <button onClick={() => setCount((count) => count + 1)}>
  //         count is {count}
  //       </button>
  //     </div>
  //   </>
  // )
}

