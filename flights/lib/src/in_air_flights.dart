import 'package:flutter/material.dart';
import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
// import 'package:cloud_firestore/cloud_firestore.dart';


class InAirFlightsView extends StatefulWidget {
    static const routeName = '/flights';

  const InAirFlightsView({super.key});

  @override
  State<InAirFlightsView> createState() => _InAirFlightsViewState();

}

class _InAirFlightsViewState extends State<InAirFlightsView> {

  final _firestore = FirebaseFirestore.instance.collection('inair');
  late GoogleMapController myMapController;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: StreamBuilder<QuerySnapshot<Map<String, dynamic>>>(
        stream: _firestore.snapshots(),

        builder: (BuildContext context, AsyncSnapshot<QuerySnapshot<Map<String, dynamic>>> snapshot) {
          if (snapshot.hasError) {
            return Text('Error: ${snapshot.error}');
          }

          if (snapshot.connectionState == ConnectionState.waiting) {
            return const CircularProgressIndicator();
          }

          return GoogleMap(
            initialCameraPosition: const CameraPosition(
              target: LatLng(37, -122),
              zoom: 6.0,
            ),
            mapType: MapType.normal,
            markers: snapshot.data!.docs.map((doc) {
              final data = doc.data();
              print(data.toString());
              final latitude = data['latitude'] as double;
              print(latitude.toString());
              final longitude = data['longitude'] as double;
              print(longitude.toString());
              final altitude = data['altitude'] as double;
              print(altitude.toString());

              return Marker(
                markerId: MarkerId(doc.id),
                position: LatLng(latitude, longitude),
                infoWindow: InfoWindow(
                  title: 'Altitude: $altitude',
                ),
              );
            }).toSet(),
            onMapCreated: (GoogleMapController controller) {
              myMapController = controller;
            },
          );

          // return Text(snapshot.data!.docs.first.data().toString());
        },
      ),
    );
  }
}
