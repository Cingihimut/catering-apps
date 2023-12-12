import { authUserSession } from "@/lib/aut-hlibs"
import Link from "next/link"

const UserActionButton = async() => {
    
    const user = await authUserSession();
    const actionLable = user? "Sign Out" : "Sign In"
    const actionUrl = user ? "/api/auth/signout" : "/api/auth/signin"
    
    return(
        <div className="flex justify-between gap-2">
            {user ? <Link className="text-color-third py-1" href="/users/dashboard"> Dashboard </Link> :null}
            <Link href={actionUrl} className="bg-color-white text-color-dark py-1 px-12 inline-block">
                {actionLable}
            </Link>
        </div>
    )
}

export default UserActionButton