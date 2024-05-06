// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/johnoppenheimer/boustifaille/database/models"

func head() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://cdn.tailwindcss.com/3.4.3?plugins=typography@0.5.12\"></script><script src=\"https://unpkg.com/maplibre-gl@latest/dist/maplibre-gl.js\"></script><link href=\"https://unpkg.com/maplibre-gl@latest/dist/maplibre-gl.css\" rel=\"stylesheet\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func loadPoints(places []models.Place) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_loadPoints_e0ec`,
		Function: `function __templ_loadPoints_e0ec(places){window._map.on("load", () => {
	console.log("loading points", places);

	window._map.addSource("places", {
		type: "geojson",
		data: {
			type: "FeatureCollection",
			features: places.map((r) => ({
				type: "Feature",
				geometry: {
					type: "Point",
					coordinates: [r.Longitude, r.Latitude],
				},
				properties: {
					title: r.Name,
				},
			})),
		},
	});

	window._map.addLayer({
		id: "places-layer",
		source: "places",
		type: "circle",
		paint: {
			"circle-radius": 10,
			"circle-color": "#007cbf",
		},
	});
});
}`,
		Call:       templ.SafeScript(`__templ_loadPoints_e0ec`, places),
		CallInline: templ.SafeScriptInline(`__templ_loadPoints_e0ec`, places),
	}
}

func loadMap() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_loadMap_1408`,
		Function: `function __templ_loadMap_1408(){const style = {
	version: 8,
	sources: {
		osm: {
			type: "raster",
			tiles: ["https://a.tile.openstreetmap.fr/osmfr/{z}/{x}/{y}.png"],
			tileSize: 256,
			attribution: "&copy; OpenStreetMap Contributors",
			maxzoom: 19,
		},
	},
	layers: [
		{
			id: "osm",
			type: "raster",
			source: "osm", // This must match the source key above
		},
	],
};

window._map = new maplibregl.Map({
	container: "map",
	style,
	center: [4.840574467303128, 45.762898305863274],
	zoom: 13,
});
}`,
		Call:       templ.SafeScript(`__templ_loadMap_1408`),
		CallInline: templ.SafeScriptInline(`__templ_loadMap_1408`),
	}
}

func Index(places []models.Place) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var3 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative w-screen h-screen\"><div id=\"map\" class=\"w-screen h-screen\"></div><div class=\"w-1/5 text-white rounded-lg p-8 bg-neutral-900 absolute left-2 top-2 bottom-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, place := range places {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><button type=\"button\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(place.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/index.templ`, Line: 81, Col: 44}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = loadMap().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = loadPoints(places).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Layout(head()).Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
