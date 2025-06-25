
import { NextResponse } from 'next/server';



export async function POST(req: Request) {
    try {
        const body = await req.json();
        const { isc_id: subCategoryId, ic_id: categoryId } = body;

        if (!subCategoryId || !categoryId) {
            return new NextResponse(JSON.stringify({ error: "Category link is required" }), { status: 400 });
        }

        const token = req.headers.get('Authorization')?.split(' ')[1];

        if (!token) {
            return new NextResponse(JSON.stringify({ error: "Authorization token is required" }), { status: 401 });
        }

        try {
            const response = await fetch(`${process.env.PV_PATH_API}/items/sub/category/link?token=${token}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body)
            });

            const data = await response.json();

            if (response.status === 201) {
                return NextResponse.json(data, { status: 201 });
            } else {
                return new NextResponse(JSON.stringify(data), { status: response.status });
            }
        } catch (error) {
            console.error('Fetch error:', error);
            return new NextResponse(JSON.stringify({ error: "Unexpected error occurred during fetch" }), { status: 500 });
        }
    } catch (error) {
        console.error('Internal server error:', error);
        return new NextResponse(JSON.stringify({ error: "Internal Server Error" }), { status: 500 });
    }
}


export async function PUT(req: Request) {
    try {
        const body = await req.json();
        const { isc_name: name, isc_description: description, isc_picture_url: pictureUrl, icl_items_category_id: categoryId } = body;

        if (!name || !description || !pictureUrl || !categoryId) {
            return new NextResponse(JSON.stringify({ error: "All field are required" }), { status: 400 });
        }

        const token = req.headers.get('Authorization')?.split(' ')[1];

        if (!token) {
            return new NextResponse(JSON.stringify({ error: "Authorization token is required" }), { status: 401 });
        }

        try {
            const response = await fetch(`${process.env.PV_PATH_API}/items/sub/category/link?token=${token}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body)
            });

            const data = await response.json();

            if (response.status === 201) {
                return NextResponse.json(data, { status: 201 });
            } else {
                return new NextResponse(JSON.stringify(data), { status: response.status });
            }
        } catch (error) {
            console.error('Fetch error:', error);
            return new NextResponse(JSON.stringify({ error: "Unexpected error occurred during fetch" }), { status: 500 });
        }
    } catch (error) {
        console.error('Internal server error:', error);
        return new NextResponse(JSON.stringify({ error: "Internal Server Error" }), { status: 500 });
    }
}




export async function GET(req: Request) {

    try {
        const token = req.headers.get('Authorization')?.split(' ')[1];

        const { searchParams } = new URL(req.url);
        const subcategoryId = searchParams.get('subcategoryId');

        if (!subcategoryId) {
            return new NextResponse(JSON.stringify({ error: "subcategoryId is required" }), { status: 400 });
        }


        if (!token || token === 'null' || token === 'undefined') {
            return new NextResponse(JSON.stringify({ error: `Authorization token is required ${token}` }), { status: 401 });
        }

        try {
            const response = await fetch(`${process.env.PV_PATH_API}/items/category/link/${subcategoryId}?token=${token}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });


            const data = await response.json();

            if (response.status === 200) {
                return NextResponse.json(data, { status: 200 });
            } else {
                return new NextResponse(JSON.stringify(data), { status: response.status });
            }
        } catch (error) {
            console.error('Fetch error:', error);
            return new NextResponse(JSON.stringify({ error: "Unexpected error occurred during fetch" }), { status: 500 });
        }
    } catch (error) {
        console.error('Internal server error:', error);
        return new NextResponse(JSON.stringify({ error: "Internal Server Error" }), { status: 500 });
    }
}
