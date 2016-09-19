package gocrc48

/*
typedef struct CRC48
{
  unsigned long crc;
}CRC48_t;

#define CRCHIGHBIT            0x800000000000L
static CRC48_t initializer = {0xecf1df57a533L};
static CRC48_t generator   = {0xeadb71093528L};
static CRC48_t complement  = {0x130edf575accL};


CRC48_t CRC48_Init()
{
 	CRC48_t ctx;
 	ctx.crc = initializer.crc;
 	return ctx;
}


void CRC48_Update(CRC48_t* ctx, const void *dataPtr, int count)
{
	unsigned char *p;
	int i;
	int bits;
  for (p = (unsigned char *)dataPtr; count; count--)
  {
    ctx->crc ^= *(p++);

    for(bits=0; bits<8; bits++)
    {
      int bit = ctx->crc & 0x01;
      ctx->crc >>= 1;
      if (bit)
      {
		ctx->crc ^= generator.crc;
      }
    }
  }
}

unsigned long CRC48_Final(const CRC48_t* ctx)
{
  CRC48_t xmit;

  xmit.crc = ctx->crc;
  xmit.crc ^= complement.crc;

  return xmit.crc;
}

unsigned long CRC48(const char* data, int count){
	CRC48_t ctx = CRC48_Init();
	CRC48_Update(&ctx, data, count);
	return CRC48_Final(&ctx);
}
*/
import "C"
import "unsafe"

func CRC48(buf []byte) int64 {
	p := (*C.char)(unsafe.Pointer(&buf[0]))
	return int64(C.CRC48(p, C.int(len(buf))))
}
