// Provides services for working with codecs.
//
// The CodecService interface defines various functions that make it much easier to obtain a
// codec object that is appropriate for the data you wish to handle.
//
// To write a new codec service, simply confrom to the CodecService interface, and install it by
// doing:
//
//    myCodec := new(MyCodec)
//    services.InstalledCodecs = append(services.InstalledCodecs, myCodec)
package services